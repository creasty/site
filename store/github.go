package store

import (
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	oauth2github "golang.org/x/oauth2/github"

	"github.com/creasty/site/utils"
)

type GithubApplicationClient struct {
	conf *oauth2.Config
}

func NewGithubApplicationClient() *GithubApplicationClient {
	conf := &oauth2.Config{
		ClientID:     utils.Config.GithubClientID,
		ClientSecret: utils.Config.GithubClientSecret,
		Scopes:       []string{},
		Endpoint:     oauth2github.Endpoint,
	}

	return &GithubApplicationClient{conf}
}

func (self *GithubApplicationClient) AuthCodeURL() string {
	return self.conf.AuthCodeURL("state", oauth2.AccessTypeOnline)
}

func (self *GithubApplicationClient) Exchange(token string) (*oauth2.Token, error) {
	return self.conf.Exchange(oauth2.NoContext, token)
}

type GithubUserClient struct {
	*github.Client
}

func NewGithubUserClient(token string) *GithubUserClient {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	client := github.NewClient(oauth2.NewClient(oauth2.NoContext, ts))

	return &GithubUserClient{client}
}

func (self *GithubUserClient) User() (user *github.User, err error) {
	user = &github.User{}

	req, err := self.Client.NewRequest("GET", "/user", nil)
	if err != nil {
		return
	}

	res, err := self.Client.Do(req, user)
	if err != nil {
		return
	}
	defer res.Body.Close()

	return
}
