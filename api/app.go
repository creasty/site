package api

import (
	log "github.com/Sirupsen/logrus"
)

type AppController struct{}

func NewAppController() *AppController {
	return &AppController{}
}

func (self *AppController) Log(args ...interface{}) {
	log.Info(args...)
}
