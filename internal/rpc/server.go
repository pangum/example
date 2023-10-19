package rpc

import (
	"github.com/pangum/example/internal/rpc/internal"
	"github.com/pangum/example/internal/rpc/internal/server"
	"github.com/pangum/grpc"
)

type Server struct {
	grpc.Register

	grpc *grpc.Server
	tag  *server.Tag
}

func newServer(get internal.Get) *Server {
	return &Server{
		grpc: get.Grpc,
		tag:  get.Tag,
	}
}

func (s *Server) Start() error {
	return s.grpc.Serve(s)
}

func (s *Server) Name() string {
	return "gRPC服务器"
}

func (s *Server) Stop() (err error) {
	s.grpc.Stop()

	return
}
