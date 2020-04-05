package routes

import (
	other "../handlers/other"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/", other.DefaultHandler).Methods("GET")

	LoadToDoRoutes(r)

	return r
}
