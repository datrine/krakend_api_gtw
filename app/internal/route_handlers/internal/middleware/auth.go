package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	jwtutils "github.com/datrine/basic_crud_with_auth/internal/route_handlers/internal/jwtUtils"
	responsedtos "github.com/datrine/basic_crud_with_auth/internal/route_handlers/internal/response_dto"
	"github.com/datrine/basic_crud_with_auth/internal/route_handlers/internal/utils"
)

func Auth(w http.ResponseWriter, r *http.Request) bool {
	fmt.Printf("Auth middleware \n")
	if r.Header == nil {
		utils.GenerateResponse(w, http.StatusUnauthorized, responsedtos.Response{
			Status:  http.StatusUnauthorized,
			Message: "Nil header.",
		})
		return false
	}
	authHeaders := r.Header["Authorization"]
	if authHeaders == nil {
		utils.GenerateResponse(w, http.StatusUnauthorized, responsedtos.Response{
			Status:  http.StatusUnauthorized,
			Message: "Nil Authorization headers.",
		})
		return false
	}
	auth := authHeaders[0]
	if auth == "" {
		utils.GenerateResponse(w, http.StatusUnauthorized, responsedtos.Response{
			Status:  http.StatusUnauthorized,
			Message: "Authorization header missing.",
		})
		return false
	}
	fmt.Printf("Auth middleware \n")
	arr := strings.Split(auth, " ")
	if arr == nil {
		utils.GenerateResponse(w, http.StatusUnauthorized, responsedtos.Response{
			Status:  http.StatusUnauthorized,
			Message: "Invalid or missing jwt",
		})
		return false
	}
	if len(arr) != 2 {
		utils.GenerateResponse(w, http.StatusUnauthorized, responsedtos.Response{
			Status:  http.StatusUnauthorized,
			Message: "Malformed jwt",
		})
		return false
	}
	tkStr := arr[1]
	payload, err := jwtutils.VerifyToken(tkStr)
	if err != nil {
		utils.GenerateResponse(w, http.StatusUnauthorized, responsedtos.Response{
			Status:  http.StatusUnauthorized,
			Message: "Failed to verify jwt",
		})
		return false
	}
	ctx := context.WithValue(r.Context(), "user", payload)
	r.WithContext(ctx)
	return true
}
