package models

import (
	"time"
)

type Attendance struct {
	ID             int
	StudentID      int
	AttendanceDate time.Time
	IsPresent      bool
	CreatedDate    time.Time
}

var attendanceRegister = make(map[string]Attendance)

func InitiateAttendance() {
	attendanceRegister = map[string]Attendance{
		"1": {
			ID: 1,
			StudentID: 1,
			AttendanceDate: time.Now(),
			IsPresent: true,
			CreatedDate: time.Now(),
		},
		"2": {
			ID: 2,
			StudentID: 2,
			AttendanceDate: time.Now(),
			IsPresent: false,
			CreatedDate: time.Now(),
		},
		"3": {
			ID: 3,
			StudentID: 3,
			AttendanceDate: time.Now(),
			IsPresent: true,
			CreatedDate: time.Now(),
		},
		"4": {
			ID: 4,
			StudentID: 4,
			AttendanceDate: time.Now(),
			IsPresent: true,
			CreatedDate: time.Now(),
		},
		"5": {
			ID: 5,
			StudentID: 5,
			AttendanceDate: time.Now(),
			IsPresent: false,
			CreatedDate: time.Now(),
		},
		"6": {
			ID: 6,
			StudentID: 6,
			AttendanceDate: time.Now(),
			IsPresent: true,
			CreatedDate: time.Now(),
		},
	}
}