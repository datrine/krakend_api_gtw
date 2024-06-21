package main

import (
	"net/http"
	"strconv"

	"github.com/datrine/basic_crud_with_auth/config"
	routehandlers "github.com/datrine/basic_crud_with_auth/internal/route_handlers"
)

func main() {
	routehandlers.AllRoutes()
	port := config.GetPort()
	println("To run on port : ", port)
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		panic(err.Error())
	}
}
