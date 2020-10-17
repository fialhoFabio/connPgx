package controller

import (
	"encoding/json"
	"github.com/fialhoFabio/go_person/model"
	"github.com/fialhoFabio/go_person/repository"
	"net/http"
	"strings"
)

func PersonController(w http.ResponseWriter, req *http.Request) {
	personRepository := repository.PersonRepository{}
	switch req.Method {
	case http.MethodGet:
		id := strings.ReplaceAll(req.URL.Path, "/person/", "")
		if id == "" {
			personModelList := personRepository.GetAll()
			res, _ := json.Marshal(personModelList)
			_, _ = w.Write(res)
		} else {
			personModel, err := personRepository.GetOne(id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			} else {
				res, _ := json.Marshal(personModel)
				_, _ = w.Write(res)
			}
		}
	case http.MethodPost:
		if req.Header.Get("Content-Type") == "application/json" {
			decoder := json.NewDecoder(req.Body)
			var personModel model.Person
			_ = decoder.Decode(&personModel)
			_, err := personRepository.Insert(personModel)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			} else {
				w.WriteHeader(http.StatusCreated)
			}
		} else {
			http.Error(w, "Content-Type must be application/json", http.StatusBadRequest)
		}
	case http.MethodPut:
		if req.Header.Get("Content-Type") == "application/json" {
			id := strings.ReplaceAll(req.URL.Path, "/person/", "")
			decoder := json.NewDecoder(req.Body)
			var personModel model.Person
			_ = decoder.Decode(&personModel)
			_, err := personRepository.Update(id, personModel)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			} else {
				w.WriteHeader(http.StatusOK)
			}
		} else {
			http.Error(w, "Content-Type must be application/json", http.StatusBadRequest)
		}
	case http.MethodDelete:
		id := strings.ReplaceAll(req.URL.Path, "/person/", "")
		if id != "" {
			_, err := personRepository.Delete(id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			} else {
				w.WriteHeader(http.StatusOK)
			}
		} else {
			http.NotFound(w, req)
		}
	default:
		http.NotFound(w, req)
	}

}
