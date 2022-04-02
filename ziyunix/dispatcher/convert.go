package dispatcher

import (
	`agent/ziyunix/controller`

	`github.com/goexl/gox`
	`github.com/pangum/logging`
	`github.com/pangum/mqtt`
	`github.com/pangum/pangu`
	`github.com/ziyunix/message/convert`
)

type (
	// Convert 转码分发器
	Convert struct {
		convert *controller.Convert
		logger  *logging.Logger

		_ gox.CannotCopy
	}

	convertIn struct {
		pangu.In

		Convert *controller.Convert
		Logger  *logging.Logger
	}
)

func newConvert(in convertIn) *Convert {
	return &Convert{
		convert: in.Convert,
		logger:  in.Logger,
	}
}

func (c *Convert) OnMessage(message *mqtt.Message) (err error) {
	task := new(convert.Task)
	if err = message.Fill(task); nil != err {
		return
	}

	// 转码
	err = c.convert.Task(task)

	return
}
