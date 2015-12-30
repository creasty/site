package server

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"

	"github.com/creasty/site/api"
	"github.com/creasty/site/store"
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

		r.GET("/me", requireAuth(api.Controller.Me.Show))
	}
}

func requireAuth(fn gin.HandlerFunc) gin.HandlerFunc {
	r := regexp.MustCompile(`^Token\s+token=(?:["']?)([^"]+)(?:["']?)$`)

	return func(c *gin.Context) {
		if h := c.Request.Header.Get("Authorization"); h != "" {
			m := r.FindStringSubmatch(h)
			if len(m) != 2 {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			user, err := store.NewUserStore().FindByGithubToken(m[1])
			if err != nil {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			c.Set("currentUser", user)
		}

		fn(c)
	}
}
