package handlers

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type StudentHandler struct {
	studSvr Server
}

func NewStudentHandler(svr Server) *StudentHandler{
	return &StudentHandler{studSvr: svr}
}
func (studHand StudentHandler) GetStudentList(w http.ResponseWriter, r *http.Request) {
	log.Println("\n In handler, GetStudentList")
	// GetStudents()
	
}

func (studHand StudentHandler) GetStudent(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id != "" {
		// GetStudentById(id)
	}
}

func (studHand StudentHandler) CreateStudent(w http.ResponseWriter, r *http.Request) {
	// studInst.CreateStudent(r.Body)
}
