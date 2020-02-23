package commons

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"usermanagement/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

func GenerateToken(user *models.User) (string, error) {
	secret := viper.GetString("secret.key")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id" : user.ID,
		"email" : user.Email,
	})

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		log.Fatal(err)
	}

	return tokenString, nil
}

func TokenVerifyMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(write http.ResponseWriter, request *http.Request) {
		authHeader := request.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")

		if len(bearerToken) == 2 {
			authToken := bearerToken[1]

			token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("error: something went wrong")
				}

				return []byte(viper.GetString("secret.key")), nil
			})

			if err != nil {
				Response(write, http.StatusUnauthorized, Message(false, err.Error()))
				return
			}

			if token.Valid {
				next.ServeHTTP(write, request)
			} else {
				Response(write, http.StatusUnauthorized, Message(false, err.Error()))
				return
			}
		} else {
			Response(write, http.StatusUnauthorized, Message(false, "Invalid token"))
			return
		}
	})
}
