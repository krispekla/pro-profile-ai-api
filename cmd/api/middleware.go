package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/krispekla/pro-profile-ai-api/config"
)

func AuthMiddleware(app *config.Application) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				app.ClientError(w, http.StatusUnauthorized)
				return
			}

			splitToken := strings.Split(authHeader, "Bearer ")
			if len(splitToken) != 2 {
				app.ClientError(w, http.StatusUnauthorized)
				return
			}
			tokenStr := splitToken[1]
			token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(app.JwtSecret), nil
			})

			if err != nil {
				app.ClientError(w, http.StatusUnauthorized)
				return
			}
			var user_id interface{}
			if claims, ok := token.Claims.(jwt.MapClaims); ok {
				fmt.Println(claims["sub"], claims["exp"])
				user_id = claims["sub"]
				if user_id == nil || user_id.(string) == "" {
					app.ClientError(w, http.StatusUnauthorized)
					return
				}
			}
			ctx := context.WithValue(r.Context(), "user", "123")
			app.InfoLog.Println("User is logged in with id: ", user_id)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
