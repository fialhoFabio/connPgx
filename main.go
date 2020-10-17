package main

import (
	"fmt"
	"github.com/fialhoFabio/go_person/controller"
	"github.com/fialhoFabio/go_person/pg_connection"
	"net/http"
)

func main() {
	pg_connection.Initialize()
	defer pg_connection.Connection().Close()

	// The both of then are needed (https://stackoverflow.com/q/64398809/12448148)
	http.HandleFunc("/person", controller.PersonController)
	http.HandleFunc("/person/", controller.PersonController)
	fmt.Println("Server running...")
	_ = http.ListenAndServe(":8090", nil)
}
