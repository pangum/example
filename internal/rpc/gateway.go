package rpc

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/itcoursee/protocol"
	"github.com/pangum/grpc"
)

func (s *Server) Gateway(_ *runtime.ServeMux, _ *[]grpc.DialOption) (ctx context.Context, hs grpc.EndpointHandlers) {
	ctx = context.Background()
	hs = grpc.EndpointHandlers{
		protocol.RegisterTagHandlerFromEndpoint,
	}

	return
}
