package internal

import (
	"github.com/pangum/example/internal/internal/get"
	"github.com/pangum/example/internal/rpc"
	"github.com/pangum/pangu"
)

type Bootstrap struct {
	pangu.Lifecycle

	rpc *rpc.Server
}

func NewBootstrap(get get.Bootstrap) pangu.Bootstrap {
	return &Bootstrap{
		rpc: get.Rpc,
	}
}

func (b *Bootstrap) Startup(app *pangu.App) error {
	return app.Add(b.rpc)
}
