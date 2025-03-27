package server

import (
	"net/http"
	"student-attendance/pkg/api_middleware"
	"student-attendance/pkg/handlers"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func routesGeneral() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(api_middleware.MyMiddleware)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Student Attendance System"))
	})

	r.Mount("/student", StudentRoutes())

	r.Mount("/attendance", AttendanceRoutes())
}

func StudentRoutes() chi.Router {
	studentRouter := chi.NewRouter()
	studentHandler := handlers.StudentHandler{}
	studentRouter.Get("/", studentHandler.GetStudentList)
	studentRouter.Post("/", studentHandler.CreateStudent)
	return studentRouter
}

func AttendanceRoutes() chi.Router {
	attnRouter := chi.NewRouter()
	attendanceHandler := handlers.AttendanceHandler{}
	attnRouter.Get("/", attendanceHandler.GetAttendance)
	attnRouter.Get("/{student_id}", attendanceHandler.GetStudentAttendance)
	attnRouter.Post("/{student_id}", attendanceHandler.CreateAttendance)
	return attnRouter
}
