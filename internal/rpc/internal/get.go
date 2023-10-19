package internal

import (
	"github.com/pangum/example/internal/rpc/internal/server"
	"github.com/pangum/grpc"
	"github.com/pangum/pangu"
)

type Get struct {
	pangu.Get

	Grpc *grpc.Server
	Tag  *server.Tag
}
