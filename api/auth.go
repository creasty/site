package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/creasty/site/store"
)

type AuthController struct {
	*AppController
}

func NewAuthController() *AuthController {
	return &AuthController{NewAppController()}
}

func (self *AuthController) Show(c *gin.Context) {
	user, err := store.NewUserStore().FindByGithubToken(c.Query("token"))
	if err == nil {
		c.JSON(http.StatusOK, user)
	} else {
		c.AbortWithError(http.StatusUnauthorized, err)
	}
}

func (self *AuthController) GetAuthCodeUrl(c *gin.Context) {
	url := store.NewGithubApplicationClient().AuthCodeURL()
	c.JSON(http.StatusOK, gin.H{"url": url})
}

func (self *AuthController) Exchange(c *gin.Context) {
	authCode := c.PostForm("code")
	if authCode == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	t, err := store.NewGithubApplicationClient().Exchange(authCode)
	if err != nil || !t.Valid() {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	c.JSON(http.StatusOK, gin.H{"accessToken": t.AccessToken, "refreshToken": t.RefreshToken, "expiry": t.Expiry})
}
