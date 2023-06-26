package server

import "github.com/dptsi-bryanagamk/go-rest-miwt-ddd/container"

func InitHttp(co *container.Container) error {
	httpServer := httpapi.Server{
		AppToken: co.Config.Server.AppToken,
	}

	return runHTTPServer(httpServer, co)
}

func InitGRPC(co *container.Container) error {
	grpcserver := grcpapi.Server{
		AppToken: co.Config.Server.AppToken,
	}

	return runGRPCServer(grpcserver, co)
}
