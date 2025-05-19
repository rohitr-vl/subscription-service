package main

import (
	"database/sql"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"subscription-service/data"
	"sync"
	"syscall"
	"time"

	"github.com/alexedwards/scs/redisstore"
	"github.com/alexedwards/scs/v2"
	"github.com/gomodule/redigo/redis"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const webPort = "80"

func main() {
	// unable to read from env file & makefile, so setting it here
	variableValue := "host=localhost port=5432 user=postgres password=super369 dbname=concurrency sslmode=disable timezone=UTC connect_timeout=5"
	variableName := "DSN"

	err := os.Setenv(variableName, variableValue)
	if err != nil {
		fmt.Println("Error setting environment variable:", err)
		return
	}
	// connect to db
	fmt.Println("Start connecting to Postgres DB")
	db := initDB()
	// db.Ping()

	// create sessions for logged in users
	session := initSession()

	// create loggers
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// create channels

	// create waitgroup
	wg := sync.WaitGroup{}

	//setup the application config
	app := Config{
		Session:  session,
		DB:       db,
		Wait:     &wg,
		InfoLog:  infoLog,
		ErrorLog: errorLog,
		Models:   data.New(db),
		ErrorChan: make(chan error),
		ErrorChanDone: make(chan bool),
	}

	//set up mail
	app.Mailer = app.CreateMail()
	go app.listenForMail()

	// listen for signals
	go app.listenForShutdown()

	// listen for errors
	go app.listenForErrors()

	// listen for web connections
	app.serve()
}

func (app *Config) listenForErrors() {
	for {
		select {
		case err := <- app.ErrorChan:
			app.ErrorLog.Println(err)
		case <- app.ErrorChanDone:
			return
		}
	}
}

func (app *Config) serve() {
	// start http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}
	app.InfoLog.Println("Starting web server...")
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func initDB() *sql.DB {
	conn := connectToDB()
	if conn == nil {
		log.Panic("Cannot connect to Postgres DB!")
	}
	return conn
}

func connectToDB() *sql.DB {
	counts := 0

	dsn, err := os.LookupEnv("DSN")
	if !err {
		log.Fatal("Did not find env var, DSN. ERR:", err)
	}
	// dsn := os.Getenv("DSN")

	fmt.Printf("Got DB connection info: %s \t %d \n", dsn, len(dsn))
	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("postgres db not ready yet")
		} else {
			log.Println("connected to postgres db")
			return connection
		}
		if counts > 10 {
			return nil
		}

		log.Println("backing off for 1sec")
		time.Sleep(1 * time.Second)
		counts++
		continue
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func initSession() *scs.SessionManager {
	// register new type to add in session
	gob.Register(data.User{})

	//set up session
	session := scs.New()
	session.Store = redisstore.New(initRedis())
	session.Lifetime = 60 * time.Minute
	session.Cookie.Persist = true
	session.Cookie.Secure = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	return session
}

func initRedis() *redis.Pool {
	redisPool := &redis.Pool{
		MaxIdle: 10,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
	return redisPool
}

// For graceful shutdown of application, application should wait until all goroutines are done processing
func (app *Config) listenForShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	app.shutdown()
	os.Exit(0)
}

func (app *Config) shutdown() {
	//perform any cleanup tasks
	app.InfoLog.Println("Performing clean up tasks ...")
	time.Sleep(3 * time.Second)

	//block until waitgroup is empty, this waitgroup is avaiable on App level
	app.Wait.Wait()

	app.Mailer.DoneChan <- true
	app.ErrorChanDone <- true

	app.InfoLog.Println("Closing channels and shutting down application ...")
	close(app.Mailer.MailerChan)
	close(app.Mailer.DoneChan)
	close(app.Mailer.ErrorChan)
	close(app.ErrorChan)
	close(app.ErrorChanDone)
}

func (app *Config) CreateMail() Mail {
	// create channels
	errorChan := make(chan error)
	mailerChan := make(chan Message, 100)
	mailerDoneChan := make(chan bool)

	m := Mail{
		Domain:      "localhost",
		Host:        "localhost",
		Port:        1025,
		Encryption:  "none",
		FromAddress: "rr@vl.com",
		FromName:    "Rohit Rathod",
		ErrorChan:   errorChan,
		MailerChan:  mailerChan,
		DoneChan:    mailerDoneChan,
		Wait:        app.Wait,
	}
	return m
}
