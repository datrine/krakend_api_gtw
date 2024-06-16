package routehandlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/datrine/basic_crud_with_auth/internal/db"
	sharedexports "github.com/datrine/basic_crud_with_auth/internal/shared_exports"
)

type RegisterSuccessResponse struct {
	Response
}

func RegisterUser(w http.ResponseWriter, req *http.Request) {
	var byts []byte
	buf := bytes.NewBuffer(byts)
	_, err := buf.ReadFrom(req.Body)
	if err != nil {
		err = generateResponse(w, &RegisterSuccessResponse{
			Response: Response{
				Status:  http.StatusBadRequest,
				Message: "Bad request.",
			},
		})
	}
	u := &sharedexports.CreateUser{}
	err = json.Unmarshal(buf.Bytes(), u)
	if err != nil {
		err = generateResponse(w, &RegisterSuccessResponse{
			Response: Response{
				Status:  http.StatusBadRequest,
				Message: "Bad request.",
			},
		})
	}
	repo := db.NewConn()
	err = repo.CreateUser(context.TODO(), u)
	if err != nil {
		err = generateResponse(w, &RegisterSuccessResponse{
			Response: Response{
				Status:  http.StatusBadRequest,
				Message: "Bad request.",
			},
		})
	}

	err = generateResponse(w, &RegisterSuccessResponse{
		Response: Response{
			Status:  http.StatusCreated,
			Message: "User registered.",
		},
	})
}
