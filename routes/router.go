package routes

import (
	"github.com/gorilla/mux"
	other "github.com/zerefwayne/go-psql-rest-jwt-docker-boilerplate/handlers/other"
)

func NewRouter() *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/", other.DefaultHandler).Methods("GET")

	LoadAuthRoutes(r)
	LoadToDoRoutes(r)

	return r
}
