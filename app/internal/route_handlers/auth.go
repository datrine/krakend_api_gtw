package routehandlers

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/datrine/basic_crud_with_auth/internal/db"
	jwtutils "github.com/datrine/basic_crud_with_auth/internal/route_handlers/internal/jwtUtils"
	responsedtos "github.com/datrine/basic_crud_with_auth/internal/route_handlers/internal/response_dto"
	"github.com/datrine/basic_crud_with_auth/internal/route_handlers/internal/utils"
	sharedexports "github.com/datrine/basic_crud_with_auth/internal/shared_exports"
	jwtV5 "github.com/golang-jwt/jwt/v5"
)

type MyCustomClaims struct {
	jwtV5.RegisteredClaims
}

func BasicLogin(w http.ResponseWriter, req *http.Request) {
	var byts []byte
	buf := bytes.NewBuffer(byts)
	_, err := buf.ReadFrom(req.Body)
	if err != nil {
		(utils.GenerateResponse(w, http.StatusBadRequest, &responsedtos.Response{
			Status:  http.StatusBadRequest,
			Message: "Bad request: " + err.Error(),
		}))
		return
	}
	exp := &sharedexports.LoginUserRequest{}
	json.Unmarshal(buf.Bytes(), exp)
	if exp.Email == "" {
		(utils.GenerateResponse(w, http.StatusBadRequest, &responsedtos.Response{
			Status:  http.StatusBadRequest,
			Message: "Bad request: email address is required",
		}))
		return
	}
	fmt.Printf("%#v\n", exp)
	repo := db.GetRepository()
	user, err := repo.QueryUserByEmail(context.TODO(), exp.Email)
	if err != nil {
		(utils.GenerateResponse(w, http.StatusBadRequest, &responsedtos.Response{
			Status:  http.StatusBadRequest,
			Message: "Bad request: " + err.Error(),
		}))
		return
	}
	hashedPWd := hex.EncodeToString(md5.New().Sum([]byte(exp.Password)))
	if hashedPWd != user.PasswordHash {
		(utils.GenerateResponse(w, http.StatusBadRequest, &responsedtos.Response{
			Status:  http.StatusBadRequest,
			Message: "Bad request: Password/email not match with record",
		}))
		return
	}
	access_token, err := jwtutils.GenerateToken(&jwtutils.TokenPayload{
		Email:  user.Email,
		UserId: user.Email,
	})
	if err != nil {
		(utils.GenerateResponse(w, http.StatusBadRequest, &responsedtos.Response{
			Status:  http.StatusBadRequest,
			Message: "Bad request: " + err.Error(),
		}))
		return
	}
	resData := &responsedtos.LoginSuccessResponseData{
		AccessToken: access_token,
		LastName:    user.LastName,
		FirstName:   user.FirstName,
		Email:       user.Email,
	}

	(utils.GenerateResponse(w, http.StatusBadRequest, &responsedtos.LoginSuccessResponse{
		Status:  http.StatusBadRequest,
		Message: "Login successful ",
		Data:    resData,
	}))
}
