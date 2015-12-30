package store

import (
	"github.com/creasty/site/model"
)

type UserStore struct{}

func NewUserStore() *UserStore {
	return &UserStore{}
}

func (self *UserStore) FindByGithubToken(token string) (user *model.User, err error) {
	user = &model.User{}

	client := NewGithubUserClient(token)
	ghUser, err := client.User()
	if err != nil {
		return
	}

	err = Database.Find(&user, model.User{Github: *ghUser.Login}).Error
	return
}
