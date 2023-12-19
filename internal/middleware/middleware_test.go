package middleware

// import (
// 	"bytes"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// )

// func TestAuthMiddleware(t *testing.T) {
// 	tests := []struct {
// 		name  string
// 		token string
// 		want  int
// 	}{
// 		{
// 			name:  "passed valid token",
// 			token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJ1c2VyX3V1aWRfMTIzNCIsIm5hbWUiOiJKb2huIERvZSIsImlhdCI6MTY2ODUyNjg1OCwiZXhwIjoyMDQ3MTMxNjU4fQ.g4Qb4ef71TupJpVvuLpkI7z7Xgp620PF_Mx-P0LClH4",
// 			want:  http.StatusOK,
// 		},
// 		{
// 			name:  "passed empty token",
// 			token: "",
// 			want:  http.StatusUnauthorized,
// 		},
// 		{
// 			name:  "passed invalid token",
// 			token: "eyJhbGciOiJIUNiIsInR5cCI6IkpXVCJ9.eyJzdiJ1c2VyXMTIzNCIsIm5hbWUiOiJKb2huIERvZSIsImlhdCI6MTUxNjIzOTAyMn0.nBXKtveFmwVfs658o-5hBM",
// 			want:  http.StatusUnauthorized,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			rr := httptest.NewRecorder()

// 			r, err := http.NewRequest("GET", "/api/packages", nil)
// 			if err != nil {
// 				t.Fatal(err)
// 			}

// 			cfg := new(Application)
// 			cfg.JwtSecret = "test123"
// 			var buf bytes.Buffer
// 			cfg.InfoLog = log.New(&buf, "", log.Lshortfile)
// 			next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 				w.WriteHeader(http.StatusOK)
// 			})
// 			if tt.token != "" {
// 				r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", tt.token))
// 			}

// 			AuthMiddleware(cfg)(next).ServeHTTP(rr, r)

// 			rs := rr.Result()

// 			if rs.StatusCode != tt.want {
// 				t.Errorf("handler returned wrong status code: got %v want %v", rs.StatusCode, tt.want)
// 			}
// 			defer rs.Body.Close()
// 		})
// 	}
// }
