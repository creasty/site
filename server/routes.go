package server

import (
	"github.com/gin-gonic/gin"

	"github.com/creasty/site/api"
)

func drawRoutes(r *gin.Engine) {
	drawApiRoutes(r.Group("/api"))
}

func drawApiRoutes(r *gin.RouterGroup) {
	r.GET("/ping", api.Controller.Ping.Index)
}
