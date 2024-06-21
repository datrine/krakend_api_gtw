package routehandlers

import (
	"net/http"

	"github.com/datrine/basic_crud_with_auth/internal/route_handlers/internal/middleware"
)

func AllRoutes() {
	http.HandleFunc("GET /users/{email}", GetUserByEmail)
	http.HandleFunc("PUT /users/{email}", middleware.Chain(middleware.Auth)(UpdateUser))
	http.HandleFunc("POST /users", RegisterUser)
	http.HandleFunc("POST /auth/login/basic", BasicLogin)
}
