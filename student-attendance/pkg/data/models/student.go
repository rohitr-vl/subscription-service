package models

import "time"

type Student struct {
	ID        int
	FullName  string
	Class     int
	Section   string
	Age       int
	CreatedAt time.Time
}

func addStudent() *Student {
	return &Student{
		ID: 1,
		FullName: "Alexander Bell",
		Class: 1,
		Section: "A",
		Age: 6,
		CreatedAt: time.Now(),
	}
}