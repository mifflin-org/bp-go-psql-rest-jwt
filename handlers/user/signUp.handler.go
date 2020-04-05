package handlers

import (
	helper "../../helpers/postgres/user"
	"../../models"
	"../../utils"
	"encoding/json"
	"errors"
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
			body["error"] = errors.New("email already exists")
			utils.Respond(w, http.StatusBadRequest, false, body)
			return
		}

		newUser := new(models.User)

		newUser.Email = requestBody.Email
		newUser.Name = requestBody.Name
		newUser.Password = requestBody.Password
		newUser.CreatedAt = time.Now()
		newUser.ID = string(time.Now().UnixNano())

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