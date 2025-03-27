package services

import (
	"net/http"
	"student-attendance/pkg/data/models"
)

type student Interface {
	getStudentById(w http.ResponseWriter, r *http.Request) (*student, error)
	getStudents(w http.ResponseWriter, r *http.Request) ([]student, error)
	createStudent(w http.ResponseWriter, r *http.Request) (*student, error)
}

type StudentInst struct {
}

func (studInst StudentInst) getStudentById(w http.ResponseWriter, r *http.Request) {

}

func (studInst StudentInst) getStudents(w http.ResponseWriter, r *http.Request) {

}

func (studInst StudentInst) createStudent(w http.ResponseWriter, r *http.Request) {

}
