package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"net/http"
	"strings"
	"usermanagement/commons"
)

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
				commons.Response(write, http.StatusUnauthorized, commons.Message(false, err.Error()))
				return
			}

			if token.Valid {
				next.ServeHTTP(write, request)
			} else {
				commons.Response(write, http.StatusUnauthorized, commons.Message(false, err.Error()))
				return
			}
		} else {
			commons.Response(write, http.StatusUnauthorized, commons.Message(false, "Invalid token"))
			return
		}
	})
}
