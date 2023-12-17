package types

type JwtUser struct {
	Id        string `json:"sub"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	SessionId string `json:"session_id"`
}

type ContextKey string

const (
	UserContextKey ContextKey = "user"
)
