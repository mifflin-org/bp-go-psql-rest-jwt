package todo

import (
	"github.com/gorilla/mux"
	helper "github.com/zerefwayne/go-postgres-rest-docker-boilerplate/helpers/postgres/todo"
	"github.com/zerefwayne/go-postgres-rest-docker-boilerplate/utils"
	"net/http"
	"strconv"
)


// FetchToDoByID	GET	/todo/{id}	fetches todo by id
func FetchToDoByID(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	if id, err := strconv.ParseInt(params["id"], 10, 64); err != nil {

		body := make(map[string]interface{})

		body["error"] = err

		utils.Respond(w, http.StatusInternalServerError, false, body)

	} else {

		if todo, err := helper.FetchByID(int(id)); err != nil {

			body := make(map[string]interface{})

			body["error"] = err

			utils.Respond(w, http.StatusInternalServerError, false, body)

		} else {

			body := make(map[string]interface{})

			body["todo"] = todo

			utils.Respond(w, http.StatusOK, true, body)

		}
	}

}