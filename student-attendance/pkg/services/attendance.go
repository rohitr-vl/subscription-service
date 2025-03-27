package services

import (
	"net/http"
	"student-attendance/pkg/data/models"
)

type Attendance interface {
	CreateAttendance(w http.ResponseWriter, r *http.Request) (*models.Student, error)
	GetAttendance(w http.ResponseWriter, r *http.Request) (*models.Student, error)
	GetStudentAttendance(w http.ResponseWriter, r *http.Request) (*models.Student, error)
}

type AttendanceInst struct {
}

// CreateAttendance implements Attendance.
func (a *AttendanceInst) CreateAttendance(w http.ResponseWriter, r *http.Request) (*models.Student, error) {
	panic("unimplemented")
}

// GetAttendance implements Attendance.
func (a *AttendanceInst) GetAttendance(w http.ResponseWriter, r *http.Request) (*models.Student, error) {
	panic("unimplemented")
}

// GetStudentAttendance implements Attendance.
func (a *AttendanceInst) GetStudentAttendance(w http.ResponseWriter, r *http.Request) (*models.Student, error) {
	panic("unimplemented")
}

func NewAttendance() Attendance {
	return &AttendanceInst{}
}
