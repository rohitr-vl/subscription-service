package dto

import "time"

type Product struct {
	ID             int       `json:"id"`
	StudentID      int       `json:"student_id"`
	AttendanceDate time.Time `json:"attendance_date"`
	IsPresent      bool      `json:"is_present"`
	CreatedDate    time.Time
}
