package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/student-api/student"
	"go.uber.org/zap"
)

func createServerHandler(service student.Service, logger *zap.Logger) (http.Handler, error) {
	router := chi.NewRouter()
	router.Route("/students", func(router chi.Router) {
		router.Put("/", errorWrapper(saveStudent(service)))
		router.Get("/{studentId}", errorWrapper(getStudent(service)))
		router.Delete("/{studentId}", errorWrapper(deleteStudent(service)))
		router.Patch("/{studentId}", errorWrapper(updateStudent(service)))
	})
	return router, nil
}
