package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/krispekla/pro-profile-ai-api/internal/handler"
	"github.com/krispekla/pro-profile-ai-api/types"
)

type Middleware struct {
	JwtSecret string
	InfoLog   *log.Logger
}

func NewMiddleware(jwtSecret string, infoLog *log.Logger) *Middleware {
	return &Middleware{
		JwtSecret: jwtSecret,
		InfoLog:   infoLog,
	}
}

func (app *Middleware) AuthMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				handler.ClientError(w, http.StatusUnauthorized)
				return
			}

			splitToken := strings.Split(authHeader, "Bearer ")
			if len(splitToken) != 2 {
				handler.ClientError(w, http.StatusUnauthorized)
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
				handler.ClientError(w, http.StatusUnauthorized)
				return
			}
			user := &types.JwtUser{}
			if claims, ok := token.Claims.(jwt.MapClaims); ok {
				if userId, ok := claims["sub"].(string); ok {
					user.Id = userId
				} else {
					handler.ClientError(w, http.StatusUnauthorized)
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
