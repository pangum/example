package config

import (
	"github.com/pangum/pangu"
)

func init() {
	pangu.New().Get().Dependencies().Build().Provide(
		config,
		tag,
	)
}
