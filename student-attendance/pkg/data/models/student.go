package models

import (
	"sync"
	"time"
)

type Student struct {
	ID        int
	FullName  string
	Class     int
	Section   string
	Age       int
	CreatedAt time.Time
}

var (
	allStudentsList = make(map[string]Student)
	mux      sync.Mutex
)

func addStudent() *Student {
	return &Student{
		ID: 10,
		FullName: "Alexander Bell",
		Class: 1,
		Section: "A",
		Age: 6,
		CreatedAt: time.Now(),
	}
}

func InitiateStudents() {
	allStudentsList = map[string]Student{
		"1": {
			ID: 1,
			FullName: "Nancy Drew",
			Class: 2,
			Section: "B",
			Age: 7,
			CreatedAt: time.Now(),
		},
		"2": {
			ID: 2,
			FullName: "Jack Doe",
			Class: 1,
			Section: "C",
			Age: 6,
			CreatedAt: time.Now(),
		},
		"3": {
			ID: 3,
			FullName: "John Smith",
			Class: 3,
			Section: "A",
			Age: 8,
			CreatedAt: time.Now(),
		},
	}
}