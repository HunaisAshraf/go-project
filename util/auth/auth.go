package jwtauth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Auth interface {
	GenerateAccessToken(_id string) (string, error)
	GenerateRefreshToken(_id string) (string, error)
	VerifyToken(token string) bool
}

type JWTAuth struct {
	accessTokenString  []byte
	refreshTokenString []byte
}

func NewJWTAuth(accessTokenString string, refreshTokenString string) *JWTAuth {
	return &JWTAuth{accessTokenString: []byte(accessTokenString), refreshTokenString: []byte(refreshTokenString)}
}

func (j *JWTAuth) GenerateAccessToken(_id string) (string, error) {
	fmt.Println("generating access token")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"_id": _id,
		"exp": time.Now().Add(time.Minute * 10).Unix(),
	})

	tokenString, err := token.SignedString(j.refreshTokenString)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (j *JWTAuth) GenerateRefreshToken(id string) (string, error) {
	fmt.Println("generating refresh token")

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

func (j *JWTAuth) VerifyToken(token string) bool {
	return true
}
