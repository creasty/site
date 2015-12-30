package store

import (
	"fmt"

	"github.com/garyburd/redigo/redis"

	"github.com/creasty/site/model"
)

type UserStore struct{}

func NewUserStore() *UserStore {
	return &UserStore{}
}

func (self *UserStore) FindByGithubToken(token string) (user *model.User, err error) {
	user = &model.User{}
	key := fmt.Sprintf("UserStore.FindByGithubToken.%s", token)

	if userId, _err := redis.Int(Redis.Get(key)); _err == nil {
		err = Database.Find(&user, userId).Error
		return
	}

	client := NewGithubUserClient(token)
	ghUser, err := client.User()
	if err != nil {
		return
	}

	err = Database.Find(&user, model.User{Github: *ghUser.Login}).Error
	if err != nil {
		return
	}

	err = Redis.Set(key, user.ID)

	return
}
