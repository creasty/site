package api

var Controller = struct {
	Ping *PingController
}{}

func InitApi() {
	Controller.Ping = NewPingController()
}
