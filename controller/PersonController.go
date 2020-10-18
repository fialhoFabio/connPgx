package controller

import (
	"github.com/fialhoFabio/go_person/helper"
	"github.com/fialhoFabio/go_person/model"
	"github.com/fialhoFabio/go_person/repository"
	"net/http"
)

func PersonController(w http.ResponseWriter, req *http.Request) {
	personRepository := repository.PersonRepository{}
	switch req.Method {
	case http.MethodGet:
		pathId := helper.PathId(req, "/person/")
		if pathId.IsNull {
			personModelList := personRepository.GetAll()
			helper.ResponseJson(w, personModelList, http.StatusOK)
		} else {
			personModel := personRepository.GetOne(pathId.Data)
			helper.ResponseJson(w, personModel, http.StatusOK)
		}

	case http.MethodPost:
		var personModel model.Person
		helper.GetBodyJson(req, &personModel)
		lastInsertId := personRepository.Insert(personModel)
		helper.ResponseJson(w, lastInsertId, http.StatusCreated)

	case http.MethodPut:
		pathId := helper.PathId(req, "/person/")
		var personModel model.Person
		helper.GetBodyJson(req, &personModel)
		personRepository.Update(pathId.Data, personModel)
		helper.ResponseJson(w, nil, http.StatusOK)

	case http.MethodDelete:
		pathId := helper.PathId(req, "/person/")
		if pathId.IsNull {
			http.Error(w, "required id on url", http.StatusBadRequest)
		} else {
			personRepository.Delete(pathId.Data)
			helper.ResponseJson(w, nil, http.StatusOK)
		}

	default:
		http.NotFound(w, req)
	}
}
