package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"os"
)

// GetJWTKey retrieves the JWT secret key from environment variables.
func GetJWTKey() []byte {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		// Default key if not set (for development purposes only)
		jwtSecret = "default_secret_key"
	}
	return []byte(jwtSecret)
}

// Claims struct to use for JWT payload
type Claims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

// GenerateJWT generates a JWT for a user.
func GenerateJWT(userID int) (string, error) {
	expirationTime := time.Now().Add(72 * time.Hour)
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtKey := GetJWTKey()
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// VerifyJWT verifies the JWT and returns the user ID if valid.
func VerifyJWT(tokenString string) (int, error) {
	claims := &Claims{}
	jwtKey := GetJWTKey()
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				return 0, errors.New("token expired or not valid yet")
			}
		}
		return 0, errors.New("failed to parse token")
	}
	if !token.Valid {
		return 0, errors.New("invalid token")
	}
	return claims.UserID, nil
}
