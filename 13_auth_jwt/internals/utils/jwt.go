package utils

import (
	"errors"
	"log/slog"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWT claims
type Claims struct {
	jwt.RegisteredClaims
	Role string `json:"role"`
}

func CreateToken(jwtSecret, userID, role string) (string, error) {

	now := time.Now().UTC()
	exp := now.Add(time.Hour * 24)

	// creating a custom claim
	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "auth-service", // who issued the token
			Subject:   userID,         // subject of the token
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(exp),
		},
		Role: role, //added the role to the jwt token as well
	}

	// creating a token with above mentioned claims with Signing Method HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // create the token

	// signing the token
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		slog.Error("Error creating token", "error", err)
		return "", errors.New("error creating token")
	}

	slog.Info("Token created", "token", tokenString)
	return tokenString, nil

}

func VerifyToken(jwtSecret, tokenString string) (*Claims, error) {

	// parsing the token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		slog.Error("Error parsing token", "error", err)
		return nil, errors.New("error parsing token")
	}

	// validating the token
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		slog.Error("Error validating token", "error", err)
		return nil, errors.New("error validating token")
	}

	return claims, nil
}
