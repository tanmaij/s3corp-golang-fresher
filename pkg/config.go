package pkg

import (
	"github.com/go-chi/jwtauth"
	"os"
)

// Define unique jwtAuth instance
var jwtAuthInstance *jwtauth.JWTAuth

// GetJWTAuth return unique jwtAuth instance
func GetJWTAuth() *jwtauth.JWTAuth {
	// 1. Checking jwtAuth instance already define yet
	// If it isn't defined yet, Create new
	if jwtAuthInstance == nil {
		return jwtauth.New("HS256", []byte(os.Getenv("SECRET_KEY")), nil)
	} else {
		return jwtAuthInstance
	}
}
