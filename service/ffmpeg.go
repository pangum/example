package service

import (
	`fmt`
	`os`
	`path/filepath`
	`strings`

	`github.com/goexl/gex`
	`github.com/goexl/gox`
	`github.com/pangum/logging`
	`github.com/pangum/pangu`
	`github.com/storezhang/media`
	`github.com/ziyunix/core`
	`github.com/ziyunix/message/convert`
	`google.golang.org/protobuf/types/known/timestamppb`
)

const ffmpegExec = `ffmpeg`

type (
	// Ffmpeg 转码
	Ffmpeg struct {
		notify *Notify

		logger *logging.Logger
		_      gox.CannotCopy
	}

	ffmpegIn struct {
		pangu.In

		Notify *Notify
		Logger *logging.Logger
	}
)

func newFfmpeg(in ffmpegIn) *Ffmpeg {
	return &Ffmpeg{
		notify: in.Notify,
		logger: in.Logger,
	}
}

func (f *Ffmpeg) Exec(home string, task *convert.Task) (err error) {
	// 改变状态为已创建
	f.notify.Notification(task.Id, core.Status_STARTED, timestamppb.Now())

	// 转码
	args := make([]interface{}, 0, 16)
	// 组装参数
	if err = f.parseArgs(home, task, &args); nil != err {
		return
	}
	// 改变状态为下载中
	f.notify.Notification(task.Id, core.Status_CONVERTING, timestamppb.Now())
	// 执行转码
	_, err = gex.Exec(ffmpegExec, gex.Args(args...), gex.Dir(home))

	return
}

func (f *Ffmpeg) parseArgs(_ string, _ *convert.Task, _ *[]interface{}) (err error) {
	// 为了不暴露产品逻辑，这里把代码逻辑给去掉了，可以自己根据需要去加代码
	return
}
