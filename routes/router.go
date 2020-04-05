package routes

import (
	"github.com/gorilla/mux"
	other "github.com/zerefwayne/go-postgres-rest-docker-boilerplate/handlers/other"
)

func NewRouter() *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/", other.DefaultHandler).Methods("GET")

	LoadToDoRoutes(r)

	return r
}
