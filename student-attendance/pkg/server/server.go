package server

import (
	"fmt"
	"net/http"
	"student-attendance/pkg/services"
)

type Server struct {
	attendanceService services.Attendance
	studentService services.StudentIntf
}

func NewServer() *Server {
	att := services.NewAttendance()
	studsvr := services.NewStudent()

	svr :=  &Server{
		attendanceService:att,
		studentService:studsvr,
	}
	return svr
}

func (s *Server)AttendanceService() services.Attendance {
	return s.attendanceService
}
func (s *Server)StudentService() services.StudentIntf {
	return s.studentService
}

func (s *Server) Start() {
	r := routesGeneral()
	fmt.Println("Server running on port 8081")
	http.ListenAndServe(":8081", r)
}