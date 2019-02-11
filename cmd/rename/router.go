package main

import (
	"net/http"

	"github.com/authentication-api/authentication"
	"github.com/go-chi/chi"
	"go.uber.org/zap"
	mgo "gopkg.in/mgo.v2"
)

func createServerHandler(service authentication.Service, logger *zap.Logger, session *mgo.Session) (http.Handler, error) {
	router := chi.NewRouter()
	router.Route("/students", func(router chi.Router) {
		router.Put("/", errorWrapper(saveStudent(service, session)))
		router.Get("/{studentID}", errorWrapper(getStudent(service, session)))
		router.Delete("/{studentID}", errorWrapper(deleteStudent(service, session)))
	})
	return router, nil
}
