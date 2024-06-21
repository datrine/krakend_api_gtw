package routehandlers

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/datrine/basic_crud_with_auth/internal/db"
	responsedtos "github.com/datrine/basic_crud_with_auth/internal/route_handlers/internal/response_dto"
	"github.com/datrine/basic_crud_with_auth/internal/route_handlers/internal/utils"
	sharedexports "github.com/datrine/basic_crud_with_auth/internal/shared_exports"
)

func RegisterUser(w http.ResponseWriter, req *http.Request) {
	var byts []byte
	buf := bytes.NewBuffer(byts)
	_, err := buf.ReadFrom(req.Body)
	if err != nil {
		w.Header().Add("Status", "401")
		err = utils.GenerateResponse(w, http.StatusBadRequest, &responsedtos.RegisterSuccessResponse{
			Response: responsedtos.Response{
				Status:  http.StatusBadRequest,
				Message: "Bad request.",
			},
		})
		return
	}
	u := &sharedexports.CreateUser{}
	err = json.Unmarshal(buf.Bytes(), u)
	if err != nil {
		err = utils.GenerateResponse(w, http.StatusBadRequest, &responsedtos.RegisterSuccessResponse{
			Response: responsedtos.Response{
				Status:  http.StatusBadRequest,
				Message: "Bad request.",
			},
		})
		return
	}
	hash := md5.New()
	u.PasswordHash = hex.EncodeToString(hash.Sum([]byte(u.Password)))
	u.FirstName = strings.ToLower(u.FirstName)
	u.LastName = strings.ToLower(u.LastName)
	repo := db.GetRepository()
	err = repo.CreateUser(context.TODO(), u)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		err = utils.GenerateResponse(w, http.StatusBadRequest, &responsedtos.RegisterSuccessResponse{
			Response: responsedtos.Response{
				Status:  http.StatusBadRequest,
				Message: "Bad request.",
			},
		})
		return
	}

	err = utils.GenerateResponse(w, 201, &responsedtos.RegisterSuccessResponse{
		Response: responsedtos.Response{
			Status:  http.StatusCreated,
			Message: "User registered.",
		},
	})
}

func GetUserByEmail(w http.ResponseWriter, req *http.Request) {
	emailId := req.PathValue("email")
	if emailId == "" {
		_ = utils.GenerateResponse(w, http.StatusBadRequest, &responsedtos.RegisterSuccessResponse{
			Response: responsedtos.Response{
				Status:  http.StatusBadRequest,
				Message: "Bad request: Email is required as param",
			},
		})
		return
	}
	repo := db.GetRepository()
	user, err := repo.QueryUserByEmail(context.TODO(), emailId)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		err = utils.GenerateResponse(w, http.StatusBadRequest, &responsedtos.RegisterSuccessResponse{
			Response: responsedtos.Response{
				Status:  http.StatusBadRequest,
				Message: "Bad request: " + err.Error(),
			},
		})
		return
	}
	if user == nil {
		utils.GenerateResponse(w, 404, &responsedtos.Response{Status: http.StatusNotFound,
			Message: "User not found",
		})
		return
	}
	err = utils.GenerateResponse(w, 200, &responsedtos.GetUserByEmailSuccessResponse{
		Response: responsedtos.Response{Status: http.StatusOK,
			Message: "User details."},
		Data: &responsedtos.GetUserByEmailSuccessResponseData{
			Email:     user.Email,
			LastName:  user.LastName,
			FirstName: user.FirstName,
		},
	})
}

func UpdateUser(w http.ResponseWriter, req *http.Request) {
	emailId := req.PathValue("email")
	if emailId == "" {
		w.Header().Add("Status", "401")
		utils.GenerateResponse(w, http.StatusBadRequest, &responsedtos.RegisterSuccessResponse{
			Response: responsedtos.Response{
				Status:  http.StatusBadRequest,
				Message: "Bad request. Email not set in params",
			},
		})
		return
	}
	var byts []byte
	buf := bytes.NewBuffer(byts)
	_, err := buf.ReadFrom(req.Body)
	if err != nil {
		w.Header().Add("Status", "401")
		err = utils.GenerateResponse(w, http.StatusBadRequest, &responsedtos.RegisterSuccessResponse{
			Response: responsedtos.Response{
				Status:  http.StatusBadRequest,
				Message: "Bad request.",
			},
		})
		return
	}
	u := &sharedexports.UpdateUser{}
	err = json.Unmarshal(buf.Bytes(), u)
	if err != nil {
		err = utils.GenerateResponse(w, http.StatusBadRequest, &responsedtos.RegisterSuccessResponse{
			Response: responsedtos.Response{
				Status:  http.StatusBadRequest,
				Message: "Bad request.",
			},
		})
		return
	}
	repo := db.GetRepository()
	err = repo.UpdateUser(context.TODO(), emailId, u)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		err = utils.GenerateResponse(w, http.StatusBadRequest, &responsedtos.RegisterSuccessResponse{
			Response: responsedtos.Response{
				Status:  http.StatusBadRequest,
				Message: "Bad request.",
			},
		})
		return
	}

	err = utils.GenerateResponse(w, http.StatusOK, &responsedtos.RegisterSuccessResponse{
		Response: responsedtos.Response{
			Status:  http.StatusOK,
			Message: "User info updated.",
		},
	})
}
