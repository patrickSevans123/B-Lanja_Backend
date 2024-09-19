package routes

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// Define your JWT secret key (use a secure key in production)
var jwtSecret = []byte("222b224fd5c59f7f304a999de07283015a97122aecfa4693b2469bbcd139ba4b")

// GenerateToken generates a new JWT token
func GenerateToken(userID uint) (string, error) {
	// Create a new token object
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign and return the token
	return token.SignedString(jwtSecret)
}

// ParseToken parses and validates a JWT token
func ParseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Make sure the token method is what you expect
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return jwtSecret, nil
	})
}
