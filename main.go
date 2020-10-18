package main

import (
	"github.com/fialhoFabio/go_person/controller"
	"github.com/fialhoFabio/go_person/middleware"
	"github.com/fialhoFabio/go_person/pg_connection"
	"log"
	"net/http"
)

func main() {
	pg_connection.Initialize()
	defer pg_connection.Connection().Close()

	http.HandleFunc("/person", middleware.Load(controller.PersonController))
	http.HandleFunc("/person/", middleware.Load(controller.PersonController))

	log.Println(http.ListenAndServe(":8090", nil))
}
