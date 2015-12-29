package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingController struct {
	*AppController
}

func NewPingController() *PingController {
	return &PingController{NewAppController()}
}

func (self *PingController) Index(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
