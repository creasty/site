package api

var Controller = struct {
	Ping *PingController
}{}

func init() {
	Controller.Ping = NewPingController()
}
