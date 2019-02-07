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
	router.Route("/users", func(router chi.Router) {
		router.Put("/", errorWrapper(addNewUser(service)))
	})
	return router, nil
}
