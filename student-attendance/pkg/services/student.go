package services

import (
	"student-attendance/pkg/data/models"
)

type StudentIntf interface {
	GetStudentById(int) (*models.Student, error)
	GetStudents() ([]models.Student, error)
	CreateStudent([]string) (*models.Student, error)
}

type StudentInst struct {
}

func (studInst *StudentInst) GetStudentById(stud_id int) (*models.Student, error) {
	panic("unimplemented")
}

func (studInst *StudentInst) GetStudents() ([]models.Student, error) {
	panic("unimplemented")
}

func (studInst *StudentInst) CreateStudent([]string) (*models.Student, error) {
	panic("unimplemented")
}

func NewStudent() StudentIntf {
	return &StudentInst{}
}