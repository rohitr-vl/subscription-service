package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"student-attendance/pkg/data/models"
	"student-attendance/pkg/services"
	"time"
)

type Server struct {
	attendanceService services.Attendance
	studentService    services.StudentIntf
	// repo
}

func NewServer() *Server {
	att := services.NewAttendance()
	studsvr := services.NewStudent()

	svr := &Server{
		attendanceService: att,
		studentService:    studsvr,
	}
	return svr
}

func (s *Server) AttendanceService() services.Attendance {
	return s.attendanceService
}
func (s *Server) StudentService() services.StudentIntf {
	return s.studentService
}

func (s *Server) Start() {
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)
	r := routesGeneral()
	port := "localhost:8081"
	srv := &http.Server{
		Addr:         port,
		Handler:      r,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// repo
	models.InitiateStudents()

	// NewSVC
	svc := &Server{}
	svc.studentService = svc.StudentService()
	svc.attendanceService = svc.AttendanceService()

	// handlers

	//routes
	routesGeneral()

	fmt.Println("Server running on port 8081")
	go func() {
		log.Println("Listening on port ", port)
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	<-stopChan
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	srv.Shutdown(ctx)
	defer cancel()
	log.Println("Server gracefully stopped!")
	// http.ListenAndServe(":8081", r)
}
