package main

import (
	"net/http"

	"github.com/authentication-api/authentication"
	"github.com/go-chi/chi"
)

func createServerHandler() (http.Handler, error) {
	config := authentication.NewConfig()
	db, err := authentication.NewMongoDB(config.MongoDBCollection)
	if err != nil {
		return nil, err
	}
	repository := authentication.NewRepository(db)
	service := authentication.NewService(repository)
	router := chi.NewRouter()
	router.Route("/students", func(router chi.Router) {
		router.Put("/", errorWrapper(addNewStudent(service)))
		router.Get("/{studentID}", errorWrapper(getStudent(service)))
	})
	return router, nil
}
