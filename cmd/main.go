package main

import (
	"net/http"

	routehandlers "github.com/datrine/basic_crud_with_auth/internal/route_handlers"
)

func main() {
	http.HandleFunc("POST /users", routehandlers.RegisterUser)
	http.ListenAndServe(":8080", nil)
}
