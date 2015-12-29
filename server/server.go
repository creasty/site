package server

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-gonic/gin"

	"github.com/creasty/site/api"
	"github.com/creasty/site/store"
	"github.com/creasty/site/utils"
)

func Run() error {
	store.InitStore()
	api.InitApi()

	servers := []*http.Server{
		newApiServer(fmt.Sprintf("%s:%d", utils.Config.Hostname, utils.Config.Port)),
	}

	return gracehttp.Serve(servers...)
}

func newApiServer(addr string) *http.Server {
	r := gin.Default()
	r.Use(recoverWrapper())
	r.Use(frontendWrapper())

	r.Static("/assets", publicPath("assets"))
	r.StaticFile("/favicon.ico", publicPath("favicon.ico"))
	r.StaticFile("/pinterest-f590c.html", publicPath("pinterest-f590c.html"))

	drawRoutes(r)

	return &http.Server{
		Addr:           addr,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func recoverWrapper() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			r := recover()

			if r == nil {
				return
			}

			err := errors.New("internal error")

			if isDevDomain(c.Request.Host) {
				switch t := r.(type) {
				case string:
					err = errors.New(t)
				case error:
					err = t
				}
			}

			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}()

		c.Next()
	}
}

func frontendWrapper() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if !strings.Contains(c.HandlerName(), "server.frontendWrapper") {
			return
		}
		if c.Request.Method != "GET" && c.Request.Method != "HEAD" {
			return
		}
		if strings.HasPrefix(c.Request.URL.Path+"/", "/api/") {
			return
		}
		if strings.HasPrefix(c.Request.URL.Path+"/", "/assets/") {
			return
		}

		c.File(publicPath("index.html"))
	}
}
