package api

var Controller = struct {
	Ping *PingController
	Auth *AuthController
}{}

func InitApi() {
	Controller.Ping = NewPingController()
	Controller.Auth = NewAuthController()
}
