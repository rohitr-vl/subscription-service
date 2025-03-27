package handlers

import "net/http"

type AttendanceHandler struct {
}

func NewAttendanceHandler(svr Server) {

}

func (attnHand AttendanceHandler) GetAttendance(w http.ResponseWriter, r *http.Request) {
	// getAttendance(w, r)
}

func (attnHand AttendanceHandler) GetStudentAttendance(w http.ResponseWriter, r *http.Request) {
	// getAttendance(w, r)
}

func (attnHand AttendanceHandler) CreateAttendance(w http.ResponseWriter, r *http.Request) {
	// createAttendance(w, r)
}
