package server

import (
	"github.com/dptsi-bryanagamk/go-rest-miwt-ddd/container"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	grpcapi "github.com/dptsi-bryanagamk/go-rest-miwt-ddd/api/grpc"
)

func runGRPCServer(grpcsrv grpcapi.Server, co *container.Container) error {
	idleConnsClosed := make(chan struct{})
	go func() {
		signals := make(chan os.Signal, 1)

		signal.Notify(signals, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
		<-signals

		grpcsrv.GracefulStop()
		log.Println("GRPC server shutdown gracefully")
		close(idleConnsClosed)
	}()

	log.Println("GRPC server running on port", co.Config.Server.GrpcPort)
	if err := grpcsrv.Serve(co); err != http.ErrServerClosed {
		return err
	}
	<-idleConnsClosed
	log.Println("GRPC server stopping")
	return nil
}
