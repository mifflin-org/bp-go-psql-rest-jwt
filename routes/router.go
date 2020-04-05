package routes

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	r := mux.NewRouter()

	LoadAuthRoutes(r)
	LoadToDoRoutes(r)

	return r
}
