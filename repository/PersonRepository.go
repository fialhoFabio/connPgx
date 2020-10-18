package repository

import (
	"github.com/fialhoFabio/go_person/model"
	"github.com/fialhoFabio/go_person/pg_connection"
)

type PersonRepository struct{}

func (receiver PersonRepository) GetOne(id uint64) model.Person {
	var personModel model.Person
	err := pg_connection.Connection().QueryRow("SELECT * FROM main.person WHERE pers_code = $1", id).Scan(&personModel.PersonCode, &personModel.PersonName)
	GetError(err)
	return personModel
}

func (receiver PersonRepository) GetAll() []model.Person {
	var personModelList []model.Person
	rows, err := pg_connection.Connection().Query("SELECT * FROM main.person")
	GetError(err)
	for rows.Next() {
		personModel := model.Person{}
		err = rows.Scan(&personModel.PersonCode, &personModel.PersonName)
		GetError(err)
		personModelList = append(personModelList, personModel)
	}
	return personModelList
}

func (receiver PersonRepository) Insert(personModel model.Person) int64 {
	var lastInsertId int64
	err := pg_connection.Connection().QueryRow("INSERT INTO main.person (pers_name) VALUES ($1) returning pers_code", personModel.PersonName).Scan(&lastInsertId)
	GetError(err)
	return lastInsertId
}

func (receiver PersonRepository) Update(id uint64, personModel model.Person) {
	_, err := pg_connection.Connection().Exec("UPDATE main.person SET pers_name = $1 WHERE pers_code = $2", personModel.PersonName, id)
	GetError(err)
}

func (receiver PersonRepository) Delete(id uint64) {
	_, err := pg_connection.Connection().Exec("DELETE FROM main.person WHERE pers_code = $1", id)
	GetError(err)
}

func GetError(err error) {
	if err != nil {
		panic(err)
	}
}
