package main

import (
	`agent/ziyunix`

	`github.com/pangum/pangu`
)

type (
	bootstrap struct {
		application *pangu.Application
		agent       *ziyunix.Agent
	}

	bootstrapIn struct {
		pangu.In

		Application *pangu.Application
		Agent       *ziyunix.Agent
	}
)

func newBootstrap(in bootstrapIn) pangu.Bootstrap {
	return &bootstrap{
		application: in.Application,
		agent:       in.Agent,
	}
}

func (b *bootstrap) Setup() error {
	return b.application.AddServes(b.agent)
}
