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
	return &models.allStudentsList[stud_id], nil
}

func (studInst *StudentInst) GetStudents() ([]models.Student, error) {
	models.allStudentsList
}

func (studInst *StudentInst) CreateStudent(postData []string) (*models.Student, error) {
	models.allStudentsList = append(models.allStudentsList, postData)
}

func NewStudent() *StudentInst {
	return &StudentInst{}
}