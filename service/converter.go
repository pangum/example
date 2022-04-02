package service

import (
	`github.com/goexl/gox`
	`github.com/pangum/logging`
	`github.com/pangum/pangu`
	`github.com/ziyunix/core`
	`github.com/ziyunix/message/convert`
	`google.golang.org/protobuf/types/known/timestamppb`
)

type (
	// Converter 转码
	Converter struct {
		ffmpeg   *Ffmpeg
		uploader *Uploader
		notify   *Notify

		logger *logging.Logger
		_      gox.CannotCopy
	}

	convertIn struct {
		pangu.In

		Ffmpeg   *Ffmpeg
		Uploader *Uploader
		Notify   *Notify
		Logger   *logging.Logger
	}
)

func newConverter(in convertIn) *Converter {
	return &Converter{
		ffmpeg:   in.Ffmpeg,
		uploader: in.Uploader,
		notify:   in.Notify,
		logger:   in.Logger,
	}
}

func (c *Converter) Task(home string, task *convert.Task) (err error) {
	// 改变最终转码状态
	defer c.notification(task, err)
	// 改变状态为已创建
	c.notify.Notification(task.Id, core.Status_STARTED, timestamppb.Now())

	// 执行转码
	if err = c.ffmpeg.Exec(home, task); nil != err {
		return
	}

	// 上传文件
	if err = c.uploader.Upload(home, task); nil != err {
		return
	}

	return
}

func (c *Converter) notification(task *convert.Task, err error) {
	if nil == err {
		// 改变状态为转码失败
		c.notify.Notification(task.Id, core.Status_COMPLETED, timestamppb.Now())
	} else {
		// 改变状态为转码失败
		c.notify.Notification(task.Id, core.Status_FAILED, timestamppb.Now())
	}
}
