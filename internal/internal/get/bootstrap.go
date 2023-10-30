package get

import (
	"github.com/pangum/example/internal/rpc"
	"github.com/pangum/pangu"
)

type Bootstrap struct {
	pangu.Get

	Rpc *rpc.Server
}
