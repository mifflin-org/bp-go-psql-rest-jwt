package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/zerefwayne/go-psql-rest-jwt-docker-boilerplate/models"
	"time"
)

// VerifyJWT takes in tokenString and returns if the token is verified
func VerifyJWT(tokenString string) (jwt.Claims, error){

	signingKey := []byte("secret123")

	if token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	}); err != nil {
		return nil, err
	} else {
		return token.Claims, nil
	}

}

func GenerateJWTString(user models.User) (string, error) {

	signingKey := []byte("secret123")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Minute * 5).Unix(),
		"id": user.ID,
	})

	tokenString, err := token.SignedString(signingKey)

	return tokenString, err
}


