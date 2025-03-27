package main

import "student-attendance/pkg/server"

func main() {
	startServer := server.NewServer()
	startServer.Start()
}