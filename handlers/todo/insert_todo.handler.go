package todo

import (
	"encoding/json"
	"github.com/zerefwayne/go-postgres-rest-docker-boilerplate/utils"
	"net/http"
	"time"

	helper "github.com/zerefwayne/go-postgres-rest-docker-boilerplate/helpers/postgres/todo"
	"github.com/zerefwayne/go-postgres-rest-docker-boilerplate/models"
)

// InsertToDoHandler POST	/todo	inserts a new ToDo
func InsertToDoHandler(w http.ResponseWriter, r *http.Request) {

	// Close the request Body after function finishes
	defer r.Body.Close()

	// Create a new insert request to read request body
	requestBody := new(insertRequest)

	// JSON decoder decodes the body into requestBody on basis of json tags in struct
	_ = json.NewDecoder(r.Body).Decode(requestBody)

	// Create a new Todo
	newToDo := new(models.Todo)
	newToDo.Content = requestBody.Content
	newToDo.CreatedAt = time.Now()

	// Use the InsertHandler to insert it into database

	// Here lies the real benefit of seperating database queries from the handler
	// In this structure, it doesn't matter which database we are using
	// Incase we need to change the database, we just have to write the helpers with same
	// function signatures and return options
	// No change has to be made in the controllers except the package with which we are importing from

	if err := helper.Insert(newToDo); err != nil {

		body := make(map[string]interface{})

		body["error"] = err

		utils.Respond(w, http.StatusInternalServerError, false, body)

	} else {

		body := make(map[string]interface{})

		body["todo"] = newToDo

		utils.Respond(w, http.StatusOK, true, body)

	}

}
