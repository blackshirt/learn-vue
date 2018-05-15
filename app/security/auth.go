package security

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

var secretKey = "randomsecuredstring"

func Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
			func(tok *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}

				return []byte(secretKey), nil
			})
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)

			log.Printf("error getting token: %s", err)
		}
		if claims, ok := token.Claims.(*jwt.MapClaims); ok && token.Valid {
			log.Printf("JWT Authenticated OK (app: %s)", claims)
			next.ServeHTTP(w, r)
		}
	})
}
