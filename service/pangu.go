package service

import (
	`github.com/pangum/pangu`
)

func init() {
	pangu.New().Dependencies(
		newFfmpeg,
		newNotify,
		newConverter,
		newUploader,
	)
}
