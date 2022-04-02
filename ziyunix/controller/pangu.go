package controller

import (
	`github.com/pangum/pangu`
)

func init() {
	pangu.New().Dependencies(
		newConvert,
	)
}
