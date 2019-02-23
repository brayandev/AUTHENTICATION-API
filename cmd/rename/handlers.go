package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/student-api/student"
)

type handlerFuncError func(w http.ResponseWriter, r *http.Request) error

func saveStudent(service student.Service) handlerFuncError {
	return func(w http.ResponseWriter, r *http.Request) error {
		var student student.Student
		err := json.NewDecoder(r.Body).Decode(&student)
		if err != nil {
			return err
		}
		studentID, err := service.SaveStudent(context.TODO(), student)
		if err != nil {
			return err
		}
		return responseWriter(w, http.StatusCreated, &postStudentCreatedResponse{studentID})
	}
}

func getStudent(service student.Service) handlerFuncError {
	return func(w http.ResponseWriter, r *http.Request) error {
		studentID := chi.URLParam(r, "studentId")
		student, err := service.GetStudent(context.TODO(), studentID)
		if err != nil {
			return err
		}
		return responseWriter(w, http.StatusOK, student)
	}
}

func deleteStudent(service student.Service) handlerFuncError {
	return func(w http.ResponseWriter, r *http.Request) error {
		studentID := chi.URLParam(r, "studentId")
		err := service.DeleteStudent(context.TODO(), studentID)
		if err != nil {
			return err
		}
		return responseWriter(w, http.StatusNoContent, nil)
	}
}

func updateStudent(service student.Service) handlerFuncError {
	return func(w http.ResponseWriter, r *http.Request) error {
		studentID := chi.URLParam(r, "studentId")
		var student *student.UpdateStudent

		dErr := parseJSON(r.Body, &student)
		if dErr != nil {
			return dErr
		}

		updateResult, uErr := service.UpdateStudent(context.TODO(), studentID, student)
		if uErr != nil {
			return uErr
		}
		return responseWriter(w, http.StatusOK, updateResult)
	}
}

func responseWriter(w http.ResponseWriter, code int, content Versionable) error {
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
			log.Println(err)
		}
	}
}

func parseJSON(reader io.ReadCloser, out interface{}) error {
	err := json.NewDecoder(reader).Decode(out)
	if err != nil {
		return err
	}
	return nil
}
