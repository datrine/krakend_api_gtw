package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GenerateResponse(w http.ResponseWriter, StatusCode int, resBody interface{}) error {

	byts, err := json.Marshal(resBody)
	if err != nil {
		fmt.Printf(err.Error())
	}
	w.WriteHeader((StatusCode))
	_, err = w.Write(byts)
	return err
}
