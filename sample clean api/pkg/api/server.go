package api

import (
	"github.com/gin-gonic/gin"
	"github.com/thnkrn/go-gin-clean-arch/pkg/api/handler"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(h *handler.UserHandler)*ServerHTTP {
	r := gin.New()
	r.Use(gin.Logger())
	r.GET("findname", h.find())
	return &ServerHTTP{engine,engine}
}
func start(s * ServerHTTP){
 s.engine.Run()
}