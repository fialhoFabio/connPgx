package repository

import (
	"context"
	"github.com/fialhoFabio/go_person/model"
	"github.com/fialhoFabio/go_person/pg_connection"
)

type RegisterNotFoundError struct{}

func (receiver RegisterNotFoundError) Error() string {
	return "Register Not Found"
}

type CannotBeCreated struct{}

func (receiver CannotBeCreated) Error() string {
	return "Cannot be created"
}

type PersonRepository struct{}

func (receiver PersonRepository) GetOne(id string) (model.Person, error) {
	var personModel model.Person
	_ = pg_connection.Connection().QueryRow(context.Background(), "SELECT * FROM main.person WHERE pers_code = $1", id).Scan(&personModel.PersonCode, &personModel.PersonName)
	var err error
	if personModel.PersonCode == 0 {
		err = RegisterNotFoundError{}
	}
	return personModel, err
}

func (receiver PersonRepository) GetAll() []model.Person {
	var personModelList []model.Person
	rows, err := pg_connection.Connection().Query(context.Background(), "SELECT * FROM main.person")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		personModel := model.Person{}
		_ = rows.Scan(&personModel.PersonCode, &personModel.PersonName)
		personModelList = append(personModelList, personModel)
	}
	return personModelList
}

func (receiver PersonRepository) Insert(personModel model.Person) (int64, error) {
	commandTag, _ := pg_connection.Connection().Exec(context.Background(), "INSERT INTO main.person (pers_name) VALUES ($1)", personModel.PersonName)
	var err error
	if commandTag.RowsAffected() == 0 {
		err = CannotBeCreated{}
	}
	return commandTag.RowsAffected(), err
}

func (receiver PersonRepository) Update(id string, personModel model.Person) (int64, error) {
	commandTag, _ := pg_connection.Connection().Exec(context.Background(), "UPDATE main.person SET pers_name = $1 WHERE pers_code = $2", personModel.PersonName, id)
	var err error
	if commandTag.RowsAffected() == 0 {
		err = RegisterNotFoundError{}
	}
	return commandTag.RowsAffected(), err
}

func (receiver PersonRepository) Delete(id string) (int64, error) {
	commandTag, _ := pg_connection.Connection().Exec(context.Background(), "DELETE FROM main.person WHERE pers_code = $1", id)
	var err error
	if commandTag.RowsAffected() == 0 {
		err = RegisterNotFoundError{}
	}
	return commandTag.RowsAffected(), err

}
