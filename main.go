package main

import (
	"fmt"
	"github.com/fialhoFabio/go_person/controller"
	"github.com/fialhoFabio/go_person/middleware"
	"github.com/fialhoFabio/go_person/pg_connection"
	"net/http"
)

func main() {
	pg_connection.Initialize()
	defer pg_connection.Connection().Close()

	http.HandleFunc("/person", middleware.Load(controller.PersonController))
	http.HandleFunc("/person/", middleware.Load(controller.PersonController))

	fmt.Println("Server running...")
	_ = http.ListenAndServe(":8090", nil)
}
