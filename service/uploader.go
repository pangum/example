package service

import (
	`os`

	`github.com/goexl/gfx`
	`github.com/goexl/gox`
	`github.com/goexl/gox/field`
	`github.com/pangum/logging`
	`github.com/pangum/pangu`
	`github.com/ziyunix/core`
	`github.com/ziyunix/message/convert`
	`google.golang.org/protobuf/types/known/timestamppb`
)

type (
	// Uploader 文件上传器
	Uploader struct {
		notify *Notify

		logger *logging.Logger
		_      gox.CannotCopy
	}

	file struct {
		path string
		key  string
	}

	uploaderIn struct {
		pangu.In

		Notify *Notify
		Logger *logging.Logger
	}
)

func newUploader(in uploaderIn) *Uploader {
	return &Uploader{
		notify: in.Notify,
		logger: in.Logger,
	}
}

func (u *Uploader) Upload(home string, task *convert.Task) (err error) {
	// 为了不暴露产品逻辑，这里把代码逻辑给去掉了，可以自己根据需要去加代码
	return
}
