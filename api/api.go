package api

var Controller = struct {
	Ping *PingController
	Auth *AuthController
	Me   *MeController
}{}

func InitApi() {
	Controller.Ping = NewPingController()
	Controller.Auth = NewAuthController()
	Controller.Me = NewMeController()
}
