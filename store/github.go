package store

import (
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	oauth2github "golang.org/x/oauth2/github"

	"github.com/creasty/site/utils"
)

type GithubAuthenticator struct {
	conf *oauth2.Config
}

func (self *GithubAuthenticator) AuthCodeURL() string {
	return self.conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
}

func (self *GithubAuthenticator) Exchange(token string) (*oauth2.Token, error) {
	return self.conf.Exchange(oauth2.NoContext, token)
}

func NewGithubClient() *github.Client {
	conf := &clientcredentials.Config{
		ClientID:     utils.Config.GithubClientID,
		ClientSecret: utils.Config.GithubClientSecret,
		Scopes:       []string{},
	}

	return github.NewClient(conf.Client(oauth2.NoContext))
}

func NewGithubClientAuthenticator() *GithubAuthenticator {
	conf := &oauth2.Config{
		ClientID:     utils.Config.GithubClientID,
		ClientSecret: utils.Config.GithubClientSecret,
		Scopes:       []string{},
		Endpoint:     oauth2github.Endpoint,
	}

	return &GithubAuthenticator{conf}
}

func NewUserGithubClient(token string) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	return github.NewClient(oauth2.NewClient(oauth2.NoContext, ts))
}
