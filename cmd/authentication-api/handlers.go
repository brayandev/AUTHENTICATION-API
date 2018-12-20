package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/authentication-api/authentication"
)

type handlerFuncError func(w http.ResponseWriter, r *http.Request) error

func addNewUser(service authentication.Service) handlerFuncError {
	return func(w http.ResponseWriter, r *http.Request) error {
		var user authentication.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			return err
		}
		userInserted, err := service.Save(context.TODO(), user)
		if err != nil {
			return err
		}
		return responseWriter(w, http.StatusCreated, userInserted)
	}
}

type versionable interface {
	Version() string
}

func responseWriter(w http.ResponseWriter, code int, content versionable) error {
	if content == nil {
		w.WriteHeader(code)
		return nil
	}

	contentType := "application/json; charset=utf-8"

	if content.Version() != "" {
		contentType = fmt.Sprintf("application/%s+json; charset=utf-8", content.Version())
	}

	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(code)

	err := json.NewEncoder(w).Encode(content)
	if err != nil {
		return err
	}
	return nil
}

func errorWrapper(fn handlerFuncError) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := fn(w, r)
		if err != nil {
			fmt.Println(err)
		}
	}
}
