package store

import (
	"github.com/creasty/site/model"
)

type UserStore struct{}

func NewUserStore() *UserStore {
	return &UserStore{}
}

func (self *UserStore) FindByGithubToken(token string) (*model.User, error) {
	client := store.NewGithubUserClient(c.Query("token"))
	if _, err := client.User(); err != nil {
	}

	return nil, nil
}
