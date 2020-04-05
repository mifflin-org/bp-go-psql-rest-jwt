package routes

import (
	other "../handlers/other"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/", other.DefaultHandler).Methods("GET")

	LoadAuthRoutes(r)
	LoadToDoRoutes(r)

	return r
}
