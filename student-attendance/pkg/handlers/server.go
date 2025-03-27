package handlers

import "student-attendance/pkg/services"

type Server interface {
	AttendanceService() services.Attendance
	StudentService() services.StudentIntf
}
