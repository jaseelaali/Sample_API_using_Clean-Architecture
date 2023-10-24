package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jaseelaali/Sample_API_using_Clean-Architecture/pkg/api/handler"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(handler handler.UserHandler) *ServerHTTP {
	r := gin.New()
	r.GET("/login", handler.Wish)
	return &ServerHTTP{
		engine: r,
	}
}

func (s *ServerHTTP)Start(){
	s.engine.Run(":8080")
}
