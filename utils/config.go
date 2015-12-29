package utils

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type ConfigEntry struct {
	Hostname string
	Port     uint
	DevPort  uint
	DBUri    string
}

var Config *ConfigEntry = &ConfigEntry{
	Hostname: "",
	Port:     5000,
	DevPort:  5001,
	DBUri:    "postgres://",
}

func LoadConfig() {
	godotenv.Load()

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

	if dbUri := os.Getenv("DB_URI"); dbUri != "" {
		Config.DBUri = dbUri
	}
}
