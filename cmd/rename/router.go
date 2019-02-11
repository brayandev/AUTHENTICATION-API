package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/student-api/student"
	"go.uber.org/zap"
	mgo "gopkg.in/mgo.v2"
)

func createServerHandler(service student.Service, logger *zap.Logger, session *mgo.Session) (http.Handler, error) {
	router := chi.NewRouter()
	router.Route("/students", func(router chi.Router) {
		router.Put("/", errorWrapper(saveStudent(service, session)))
		router.Get("/{studentId}", errorWrapper(getStudent(service, session)))
		router.Delete("/{studentId}", errorWrapper(deleteStudent(service, session)))
	})
	return router, nil
}
