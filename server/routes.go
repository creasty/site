package server

import (
	"github.com/gin-gonic/gin"

	"github.com/creasty/site/api"
)

func drawRoutes(r *gin.Engine) {
	{
		r.Static("/assets", publicPath("assets"))
		r.StaticFile("/favicon.ico", publicPath("favicon.ico"))
		r.StaticFile("/pinterest-f590c.html", publicPath("pinterest-f590c.html"))
	}

	{
		r := r.Group("/api")

		r.GET("/ping", api.Controller.Ping.Index)
		r.GET("/auth/url", api.Controller.Auth.GetAuthCodeUrl)
		r.POST("/auth/exchange", api.Controller.Auth.Exchange)
	}
}
