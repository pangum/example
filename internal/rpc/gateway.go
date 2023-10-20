package rpc

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/itcoursee/protocol"
	"github.com/pangum/grpc"
)

func (s *Server) Gateway(_ *runtime.ServeMux, _ *[]grpc.DialOption) (ctx context.Context, handlers grpc.Handlers) {
	ctx = context.Background()
	handlers = grpc.Handlers{
		protocol.RegisterTagHandlerFromEndpoint,
	}

	return
}
