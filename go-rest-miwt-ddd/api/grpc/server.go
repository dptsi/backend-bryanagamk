package grpc

import (
	"github.com/dptsi-bryanagamk/go-rest-miwt-ddd/container"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedRiwayatCutiServiceServer
	server             *grpc.Server
	RiwayatCutiService riwayatCuti.Service
	AppToken           string
}

func (s *Server) Serve(co *container.Container) error {
	s.server = grpc.NewServer{
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_prometheus.StreamServerInterceptor,
			grpc_trace.StreamServerInterceptor(grpc_trace.WithServiceName("")),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_prometheus.StreamServerInterceptor,
			grpc_trace.UnaryServerInterceptor(grpc_trace.WithServiceName("")),
		)),
	}
}
