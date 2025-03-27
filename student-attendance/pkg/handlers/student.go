package handlers

import (
	"net/http"
	"student-attendance/pkg/services"

	"github.com/go-chi/chi"
)

type StudentHandler struct {
}

func (studHand StudentHandler) GetStudentList(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id != "" {
		services.getStudentById(w, r)
	} else {
		services.getStudents(w, r)
	}
}

func (studHand StudentHandler) CreateStudent(w http.ResponseWriter, r *http.Request) {
	services.createStudent(w, r)
}
