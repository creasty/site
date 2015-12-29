package server

import (
	"path"
	"strings"
)

const SERVER_PUBLIC_PATH = "./web/public"

func publicPath(paths ...string) string {
	paths = append([]string{SERVER_PUBLIC_PATH}, paths...)
	return path.Join(paths...)
}

func isDevDomain(host string) bool {
	isDevDomain := strings.Index(host, "localhost") == 0 ||
		strings.Index(host, "dockerhost") == 0 ||
		strings.Index(host, "test") == 0 ||
		strings.Index(host, "127.0.") == 0 ||
		strings.Index(host, "192.168.") == 0 ||
		strings.Index(host, "10.") == 0 ||
		strings.Index(host, "176.") == 0

	return isDevDomain
}
