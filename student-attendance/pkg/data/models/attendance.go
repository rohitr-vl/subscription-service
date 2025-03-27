package models

import "time"

type Attendance struct {
	ID             int
	StudentID      int
	AttendanceDate time.Time
	IsPresent      bool
	CreatedDate    time.Time
}
