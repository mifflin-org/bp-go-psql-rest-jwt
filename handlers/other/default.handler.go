package handlers

import (
	"github.com/zerefwayne/go-psql-rest-jwt-docker-boilerplate/utils"
	"net/http"
)

// DefaultHandler GET Returns a generic hello message
func DefaultHandler(w http.ResponseWriter, r *http.Request) {

	body := make(map[string]interface{})

	body["message"] = "Welcome to Golang API!"

	utils.Respond(w, http.StatusOK, true, body)

}
