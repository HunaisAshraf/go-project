package jwtauth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Auth interface {
	GenerateAccessToken(_id string) (string, error)
	GenerateRefreshToken(_id string) (string, error)
	// VerifyToken(next http.Handler) http.HandlerFunc
}

type JWTAuth struct {
	accessTokenString  []byte
	refreshTokenString []byte
}

func NewJWTAuth(accessTokenString string, refreshTokenString string) *JWTAuth {
	fmt.Println(refreshTokenString)
	return &JWTAuth{accessTokenString: []byte(accessTokenString), refreshTokenString: []byte(refreshTokenString)}
}

func (j *JWTAuth) GenerateAccessToken(_id string) (string, error) {
	fmt.Println("generating access token")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"_id": _id,
		"exp": time.Now().Add(time.Minute * 40).Unix(),
	})
	tokenString, err := token.SignedString(j.accessTokenString)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (j *JWTAuth) GenerateRefreshToken(id string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"_id": id,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString(j.refreshTokenString)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}
