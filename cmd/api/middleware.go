package main

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/krispekla/pro-profile-ai-api/config"
)

func AuthMidlleware(app *config.Application) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := ""
		// tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6InB1cmJqcHphbHhoaHZ4cWlkZW5rIiwicm9sZSI6ImFub24iLCJpYXQiOjE2OTk1MjUwNzcsImV4cCI6MjAxNTEwMTA3N30.17F_UCKXeJsq7VlG63F5jN8GNRVcvZ0y843-79HLlRQ"
		// Parse takes the token string and a function for looking up the key. The latter is especially
		// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
		// head of the token to identify which key to use, but the parsed token (head and claims) is provided
		// to the callback, providing flexibility.
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte(""), nil
		})

		if err != nil {
			app.ClientError(w, http.StatusUnauthorized)
			return
		}
		var user_id interface{}
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			fmt.Println(claims["sub"], claims["exp"])
			app.ClientError(w, http.StatusUnauthorized)
			user_id = claims["suba"]
			if user_id == nil || user_id.(string) == "" {
				app.ClientError(w, http.StatusUnauthorized)
				return
			}
		}

		// ctx := context.WithValue(r.Context(), "user", "123")

		// next.ServeHTTP(w, r.WithContext(ctx))
	})
}
