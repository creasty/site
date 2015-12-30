package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/creasty/site/model"
)

type MeController struct {
	*AppController
}

func NewMeController() *MeController {
	return &MeController{NewAppController()}
}

func (self *MeController) Show(c *gin.Context) {
	user := c.MustGet("currentUser").(*model.User)
	c.JSON(http.StatusOK, user)
}
