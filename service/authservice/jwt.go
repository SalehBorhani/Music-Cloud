package authservice

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const (
	AccessExpirationTime  time.Duration = 1 * time.Hour
	RefreshExpirationTime time.Duration = 24 * time.Hour
)

type Config struct {
	SignKey        string `koanf:"sign_key"`
	AccessSubject  string `koanf:"access_subject"`
	RefreshSubject string `koanf:"refresh_subject"`
}

type Service struct {
	config Config
}

func New(config Config) Service {
	return Service{config: config}
}

func (s Service) CreateAccessToken(id uint) (string, error) {
	claim := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(AccessExpirationTime)),
			Subject:   s.config.AccessSubject,
		},
		UserID: id,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenString, err := token.SignedString([]byte(s.config.SignKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s Service) CreateRefreshToken(id uint) (string, error) {
	claim := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(RefreshExpirationTime)),
			Subject:   s.config.RefreshSubject,
		},
		UserID: id,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenString, err := token.SignedString([]byte(s.config.SignKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s Service) VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return s.config.SignKey, nil
	})

	if err != nil || !token.Valid {
		return jwt.MapClaims{}, fmt.Errorf("invalid token: %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	}

	return nil, fmt.Errorf("could not parse claims")
}
