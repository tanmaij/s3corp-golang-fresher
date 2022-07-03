package pkg

import (
	"github.com/go-chi/jwtauth"
	"os"
)

// Define unique jwtAuth instance
var jwtAuthInstance *jwtauth.JWTAuth

// GetJWTAuth return unique jwtAuth instance
func GetJWTAuth() *jwtauth.JWTAuth {
	return jwtAuthInstance
}
func InitJWT() {
	jwtAuthInstance = jwtauth.New("HS256", []byte(os.Getenv("SECRET_KEY")), nil)
}
