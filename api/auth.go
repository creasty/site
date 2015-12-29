package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	*AppController
}

func NewAuthController() *AuthController {
	return &AuthController{NewAppController()}
}

func (self *AuthController) Index(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
