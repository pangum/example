package internal

import (
	"github.com/pangum/example/internal/internal/get"
	"github.com/pangum/example/internal/rpc"

	"github.com/pangum/pangu"
)

type Bootstrap struct {
	pangu.Lifecycle

	app *pangu.App
	rpc *rpc.Server
}

func NewBootstrap(get get.Bootstrap) pangu.Bootstrap {
	return &Bootstrap{
		app: get.App,
		rpc: get.Rpc,
	}
}

func (b *Bootstrap) Startup() error {
	return b.app.Add(b.rpc)
}
