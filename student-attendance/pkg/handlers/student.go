package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
)

type StudentHandler struct {
}

func newStudentHandler(svr Server) {

}
func (studHand StudentHandler) GetStudentList(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id != "" {
		// studInst.GetStudentById(id)
	} else {
		// studInst.GetStudents()
	}
}

func (studHand StudentHandler) CreateStudent(w http.ResponseWriter, r *http.Request) {
	// studInst.CreateStudent(r.Body)
}
