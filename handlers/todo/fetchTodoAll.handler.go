package todo

import (
	helper "github.com/zerefwayne/go-psql-rest-jwt-docker-boilerplate/helpers/postgres/todo"
	"github.com/zerefwayne/go-psql-rest-jwt-docker-boilerplate/utils"
	"net/http"
)


// InsertToDoHandler GET	/todos	fetches all todos
func FetchToDoAll(w http.ResponseWriter, r *http.Request) {

	if todos, err := helper.FetchAll(); err != nil {

		body := make(map[string]interface{})

		body["error"] = err

		utils.Respond(w, http.StatusInternalServerError, false, body)

	} else {

		body := make(map[string]interface{})

		body["todos"] = todos
		body["count"] = len(todos)

		utils.Respond(w, http.StatusOK, true, body)

	}

}