package handlers

import (
	"log"
	"net/http"
)

type AttendanceHandler struct {
	attnSvr Server
}

func NewAttendanceHandler(svr Server) *AttendanceHandler {
	return &AttendanceHandler{attnSvr: svr}
}

func (attnHand AttendanceHandler) GetAttendance(w http.ResponseWriter, r *http.Request) {
	log.Println("\n In handler, GetAttendance")
	// getAttendance(w, r)
}

func (attnHand AttendanceHandler) GetStudentAttendance(w http.ResponseWriter, r *http.Request) {
	// getAttendance(w, r)
}

func (attnHand AttendanceHandler) CreateAttendance(w http.ResponseWriter, r *http.Request) {
	// createAttendance(w, r)
}
