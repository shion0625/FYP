package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	Engine *gin.Engine
}

func NewServerHTTP() *ServerHTTP {

	engine := gin.New()

	engine.Use(gin.Logger())

	return &ServerHTTP{Engine: engine}
}

func (s *ServerHTTP) Start() error {

	return s.Engine.Run(":8000")
}
