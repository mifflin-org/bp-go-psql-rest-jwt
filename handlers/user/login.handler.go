package handlers

import (
	"encoding/json"
	helper "github.com/zerefwayne/go-psql-rest-jwt-docker-boilerplate/helpers/postgres/user"
	"github.com/zerefwayne/go-psql-rest-jwt-docker-boilerplate/utils"
	"net/http"
)

type signInRequest struct {
	Email string `json:"email" db:"email" bson:"email"`
	Password string `json:"password" db:"password" bson:"password"`
}

func SignInHandler(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	requestBody := new(signInRequest)

	_ = json.NewDecoder(r.Body).Decode(requestBody)


	if user, exists, err := helper.FetchUserByEmail(requestBody.Email); err != nil {
		body := make(map[string]interface{})
		body["error"] = err.Error()
		utils.Respond(w, http.StatusInternalServerError, false, body)
	} else {

		if exists {
			if user.Password == requestBody.Password {

				if tokenString, tokenError := utils.GenerateJWTString(user); tokenError != nil {
					body := make(map[string]interface{})
					body["error"] = tokenError.Error()
					utils.Respond(w, http.StatusUnauthorized, false, body)
				} else {
					body := make(map[string]interface{})
					body["token"] = tokenString
					utils.Respond(w, http.StatusOK, true, body)
				}
			} else {
				body := make(map[string]interface{})
				body["error"] = "authentication error: wrong password"
				utils.Respond(w, http.StatusUnauthorized, false, body)
			}

		} else {
			body := make(map[string]interface{})
			body["error"] = "authentication error: email doesn't exist"
			utils.Respond(w, http.StatusBadRequest, false, body)
		}
	}

}
