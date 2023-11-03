package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/shion0625/FYP/backend/pkg/service/token"
)

type Middleware interface {
	AuthenticateUser() gin.HandlerFunc
	AuthenticateAdmin() gin.HandlerFunc
	TrimSpaces() gin.HandlerFunc
}

type middleware struct {
	tokenService token.TokenService
}

func NewMiddleware(tokenService token.TokenService) Middleware {
	return &middleware{
		tokenService: tokenService,
	}
}
