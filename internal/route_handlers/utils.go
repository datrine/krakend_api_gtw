package routehandlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func generateResponse(w http.ResponseWriter, resBody interface{}) error {

	byts, err := json.Marshal(resBody)
	if err != nil {
		fmt.Printf(err.Error())
	}
	_, err = w.Write(byts)
	return err
}
