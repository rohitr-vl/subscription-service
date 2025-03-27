package dto

import "time"

type Student struct {
	ID        int    `json:"id"`
	FullName  string `json:"full_name"`
	Class     int    `json:"class"`
	Section   string `json:"section"`
	Age       int    `json:"age"`
	CreatedAt time.Time
}
