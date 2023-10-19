package rpc

import (
	"github.com/itcoursee/protocol"
	"google.golang.org/grpc"
)

func (s *Server) Grpc(server *grpc.Server) {
	protocol.RegisterTagServer(server, s.tag)
}
