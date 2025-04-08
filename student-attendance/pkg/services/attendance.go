package services

import (
	"log"
	"net/http"
	"student-attendance/pkg/data/models"
)

type Attendance interface {
	CreateAttendance(w http.ResponseWriter, r *http.Request) (*models.Attendance, error)
	GetAttendance(w http.ResponseWriter, r *http.Request) (*models.Attendance, error)
	GetStudentAttendance(w http.ResponseWriter, r *http.Request) (*models.Attendance, error)
}

type AttendanceInst struct {
}

// CreateAttendance implements Attendance.
func (a *AttendanceInst) CreateAttendance(w http.ResponseWriter, r *http.Request) (*models.Attendance, error) {
	log.Println("\n In service, CreateAttendance")
	return nil, nil
}

// GetAttendance implements Attendance.
func (a *AttendanceInst) GetAttendance(w http.ResponseWriter, r *http.Request) (*models.Attendance, error) {
	log.Println("\n In service, GetAttendance")
	return nil, nil
}

// GetStudentAttendance implements Attendance.
func (a *AttendanceInst) GetStudentAttendance(w http.ResponseWriter, r *http.Request) (*models.Attendance, error) {
	log.Println("\n In service, GetStudentAttendance")
	return nil, nil
}

func NewAttendance() Attendance {
	return &AttendanceInst{}
}
