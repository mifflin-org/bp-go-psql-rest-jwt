package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/zerefwayne/go-psql-rest-jwt-docker-boilerplate/utils"
	"log"
	"net/http"
	"strings"
)

func JWTMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")

		if len(tokenString) == 0 {
			body := make(map[string]interface{})
			body["error"] = "authentication error: login required"
			utils.Respond(w, http.StatusUnauthorized, false, body)
			return
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		if claims, err := utils.VerifyJWT(tokenString); err != nil {

			body := make(map[string]interface{})
			body["error"] = "authentication error: login failed, please retry."
			utils.Respond(w, http.StatusUnauthorized, false, body)

		} else {

			log.Printf("%v\n", claims.(jwt.MapClaims))
			log.Printf("%s\n", claims.(jwt.MapClaims)["id"])

			claimsMap := claims.(jwt.MapClaims)

			id := fmt.Sprintf("%v", claimsMap["id"])

			r.Header.Set("id", id)

			next.ServeHTTP(w, r)

		}

	})

}