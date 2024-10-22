package utils

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Email string
	Role  string
	jwt.StandardClaims
}

// GenerateToken will generate token for 5 hours with given data
func GenerateToken(key, email string) (string, error) {
	expTime := time.Now().Add(time.Hour * 10).Unix()

	claims := &Claims{
		Email: email,
		Role:  "admin",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime,
			Subject:   email,
			IssuedAt:  time.Now().Unix(),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := jwtToken.SignedString([]byte(key))
	if err != nil {
		log.Printf("unable to generate token for admin %v, err: %v", email, err.Error())
		return "", err
	}
	return signedToken, nil
}
