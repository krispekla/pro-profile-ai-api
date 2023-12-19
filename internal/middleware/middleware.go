package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/cors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/krispekla/pro-profile-ai-api/types"
)

type Middleware struct {
	JwtSecret   string
	InfoLog     *log.Logger
	ClientError func(w http.ResponseWriter, status int)
}

func NewMiddleware(jwtSecret string, infoLog *log.Logger, clientError func(w http.ResponseWriter, status int)) *Middleware {
	return &Middleware{
		JwtSecret:   jwtSecret,
		InfoLog:     infoLog,
		ClientError: clientError,
	}
}

func (app *Middleware) AuthMiddleware() func(next http.Handler) http.Handler {
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

			if err != nil || token == nil || !token.Valid {
				app.ClientError(w, http.StatusUnauthorized)
				return
			}
			user := &types.JwtUser{}
			if claims, ok := token.Claims.(jwt.MapClaims); ok {
				if userId, ok := claims["sub"].(string); ok {
					user.Id = userId
				} else {
					app.ClientError(w, http.StatusUnauthorized)
					return
				}
				if email, ok := claims["email"].(string); ok {
					user.Email = email
				}
				if role, ok := claims["role"].(string); ok {
					user.Role = role
				}
				if sessionId, ok := claims["session_id"].(string); ok {
					user.SessionId = sessionId
				}
			}
			ctx := context.WithValue(r.Context(), types.UserContextKey, user)
			app.InfoLog.Printf("User is logged in with id: %s", user.Id)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func (app *Middleware) CorsMiddleware() func(next http.Handler) http.Handler {
	return cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
}
