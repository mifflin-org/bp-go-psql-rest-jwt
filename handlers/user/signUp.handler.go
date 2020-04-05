package handlers

import (
	"encoding/json"
	"github.com/google/uuid"
	helper "github.com/zerefwayne/go-psql-rest-jwt-docker-boilerplate/helpers/postgres/user"
	"github.com/zerefwayne/go-psql-rest-jwt-docker-boilerplate/models"
	"github.com/zerefwayne/go-psql-rest-jwt-docker-boilerplate/utils"
	"net/http"
	"time"
)

type signUpRequest struct {
	Name string `json:"name" db:"name" bson:"name"`
	Email string `json:"email" db:"email" bson:"email"`
	Password string `json:"password" db:"password" bson:"password"`
}

func SignUpHandler(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	requestBody := new(signUpRequest)

	_ = json.NewDecoder(r.Body).Decode(requestBody)

	if exists, err := helper.DoesEmailExist(requestBody.Email); err != nil {
		body := make(map[string]interface{})
		body["error"] = err
		utils.Respond(w, http.StatusInternalServerError, false, body)
	} else {

		if exists {
			body := make(map[string]interface{})
			body["error"] = "email already exists"
			utils.Respond(w, http.StatusBadRequest, false, body)
			return
		}

		newUser := new(models.User)

		newUser.Email = requestBody.Email
		newUser.Name = requestBody.Name

		if hash, err := utils.HashPassword(requestBody.Password); err != nil {
			body := make(map[string]interface{})
			body["error"] = "error while hashing password"
			utils.Respond(w, http.StatusBadRequest, false, body)
		} else {
			newUser.Password = hash
		}

		newUser.CreatedAt = time.Now()
		newUser.ID = uuid.New().String()


		if insertError := helper.InsertUser(newUser); insertError != nil {
			body := make(map[string]interface{})
			body["error"] = err
			utils.Respond(w, http.StatusInternalServerError, false, body)
		} else {

			body := make(map[string]interface{})
			body["newUser"] = newUser
			utils.Respond(w, http.StatusOK, true, body)

		}


	}


}