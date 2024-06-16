package routehandlers

import (
	"bytes"
	"net/http"

	jwtutils "github.com/datrine/basic_crud_with_auth/internal/jwtUtils"
	jwtV5 "github.com/golang-jwt/jwt/v5"
)

type LoginSuccessResponseData struct {
	AccessToken string `json:"access_token"`
}
type LoginSuccessResponse struct {
	Status  int
	Message string
	Data    *LoginSuccessResponseData `json:"data"`
}

type MyCustomClaims struct {
	jwtV5.RegisteredClaims
}

func BasicLogin(w http.ResponseWriter, req *http.Request) {
	var byts []byte
	buf := bytes.NewBuffer(byts)
	_, err := buf.ReadFrom(req.Body)
	if err != nil {
		(generateResponse(w, &LoginSuccessResponse{}))
	}

	jwtutils.GenerateToken(&jwtutils.TokenPayload{})
}
