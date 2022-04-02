package controller

import (
	`agent/conf`
	`agent/service`

	`github.com/goexl/gox`
	`github.com/pangum/pangu`
	`github.com/ziyunix/message/convert`
)

type (
	// Convert 转码
	Convert struct {
		agent     conf.Agent
		converter *service.Converter
		_         gox.CannotCopy
	}

	convertIn struct {
		pangu.In

		Agent     conf.Agent
		Converter *service.Converter
	}
)

func newConvert(in convertIn) *Convert {
	return &Convert{
		converter: in.Converter,
	}
}

func (c *Convert) Task(task *convert.Task) error {
	return c.converter.Task(c.agent.Home, task)
}
