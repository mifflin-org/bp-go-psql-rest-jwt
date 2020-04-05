package todo

import (
	helper "../../helpers/postgres/todo"
	"../../utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)


// DeleteToDoByID	DELETE	/todo/{id}	deletes todo by id
func DeleteToDoByID(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	if id, err := strconv.ParseInt(params["id"], 10, 64); err != nil {

		body := make(map[string]interface{})

		body["error"] = err

		utils.Respond(w, http.StatusInternalServerError, false, body)

	} else {

		if err := helper.DeleteByID(int(id)); err != nil {

			body := make(map[string]interface{})

			body["error"] = err

			utils.Respond(w, http.StatusInternalServerError, false, body)

		} else {

			body := make(map[string]interface{})

			utils.Respond(w, http.StatusOK, true, body)

		}
	}

}