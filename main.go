package main

import (
	"github.com/pangum/example/internal"

	"github.com/pangum/pangu"
)

func main() {
	app := pangu.New()
	app = app.Banner().Ascii("example").Build()
	app.Get().Run(internal.NewBootstrap)
}
