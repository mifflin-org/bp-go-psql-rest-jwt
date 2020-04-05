package utils

import (
	"encoding/json"
	"net/http"
)

func Respond(w http.ResponseWriter, code int, success bool, body map[string]interface{}) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	response := new(response)

	response.Success = success
	response.Payload = body

	responseStr, _ := json.Marshal(response)

	_, _ = w.Write(responseStr)

}
