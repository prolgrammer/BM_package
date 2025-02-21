package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/prolgrammer/BM_package/jwt"
)

type middleware struct {
	manager jwt.TokenService
}

type Middleware interface {
	HandleErrors(c *gin.Context)
	Authenticate(c *gin.Context)
}

func NewMiddleware(signSecretToken string) Middleware {
	return &middleware{
		manager: jwt.NewTokenService(signSecretToken),
	}
}
