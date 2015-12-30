package server

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"

	"github.com/creasty/site/api"
	"github.com/creasty/site/store"
	"github.com/creasty/site/utils"
)

func Run() error {
	store.InitStore()
	defer store.Close()

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
	r.Use(corsWrapper())

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
		if c.Request.Method != "GET" && c.Request.Method != "HEAD" {
			c.Next()
			return
		}
		if isUnderPath(c.Request.URL.Path, "/api") {
			c.Next()
			return
		}
		if isUnderPath(c.Request.URL.Path, "/assets") {
			c.Next()
			return
		}

		c.File(publicPath("index.html"))
	}
}

func corsWrapper() gin.HandlerFunc {
	mw := cors.New(cors.Options{
		AllowedOrigins:   utils.Config.CorsOrigins,
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		Debug:            false,
	})

	return func(c *gin.Context) {
		if !isUnderPath(c.Request.URL.Path, "/api") {
			c.Next()
			return
		}
		if c.Request.Method != "OPTIONS" {
			c.Next()
			return
		}

		mw.HandlerFunc(c.Writer, c.Request)
		c.AbortWithStatus(http.StatusOK)
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
