package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	errors2 "github.com/prolgrammer/BM_package/errors"
)

type tokenService struct {
	signSecretToken []byte
}

type TokenService interface {
	Create(claims map[string]interface{}) (string, error)
	Parse(token string) (map[string]interface{}, error)
}

func NewTokenService(signSecretToken string) TokenService {
	return &tokenService{
		signSecretToken: []byte(signSecretToken),
	}
}

func (t tokenService) Create(claims map[string]interface{}) (string, error) {
	var mapClaims jwt.MapClaims
	mapClaims = claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)
	tokenString, err := token.SignedString(t.signSecretToken)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (t tokenService) Parse(token string) (map[string]interface{}, error) {
	claims := jwt.MapClaims{}
	parsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return t.signSecretToken, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors2.ErrExpired
		}
		return nil, errors.Join(errors2.ErrInvalid, err)
	}

	claims = parsed.Claims.(jwt.MapClaims)

	return claims, err
}
