package utils

import (
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type ConfigEntry struct {
	Env                string
	Hostname           string
	Port               uint
	DevPort            uint
	DatabaseUrl        string
	RedisUrl           string
	CorsOrigins        []string
	GithubClientID     string
	GithubClientSecret string
}

var Config *ConfigEntry = &ConfigEntry{
	Env:                "development",
	Hostname:           "",
	Port:               5000,
	DevPort:            5001,
	DatabaseUrl:        "",
	RedisUrl:           "",
	CorsOrigins:        []string{},
	GithubClientID:     "",
	GithubClientSecret: "",
}

func LoadConfig() {
	godotenv.Load()

	if serverEnv := os.Getenv("SERVER_ENV"); serverEnv != "" {
		Config.Env = serverEnv
	}

	if hostname := os.Getenv("SERVER_HOST"); hostname != "" {
		Config.Hostname = hostname
	}

	if port := os.Getenv("PORT"); port != "" {
		if iport, err := strconv.Atoi(port); err == nil {
			Config.Port = uint(iport)
		}
	}

	if port := os.Getenv("DEV_PORT"); port != "" {
		if iport, err := strconv.Atoi(port); err == nil {
			Config.DevPort = uint(iport)
		}
	}

	if databaseUrl := os.Getenv("DATABASE_URL"); databaseUrl != "" {
		Config.DatabaseUrl = databaseUrl
	}

	if redisUrl := os.Getenv("REDIS_URL"); redisUrl != "" {
		Config.RedisUrl = redisUrl
	}

	if corsOrigins := os.Getenv("CORS_ORIGINS"); corsOrigins != "" {
		Config.CorsOrigins = strings.Split(corsOrigins, ",")
	}

	if githubClientID := os.Getenv("GITHUB_CLIENT_ID"); githubClientID != "" {
		Config.GithubClientID = githubClientID
	}

	if githubClientSecret := os.Getenv("GITHUB_CLIENT_SECRET"); githubClientSecret != "" {
		Config.GithubClientSecret = githubClientSecret
	}
}
