package config

import (
	"github.com/pangum/pangu"
)

// Config 配置外层包装
type Config struct {
	// 程序主体配置
	Tag Tag `json:"tag" yaml:"tag" xml:"tag" toml:"tag" validate:"required"`
}

func config(config *pangu.Config) (conf *Config, err error) {
	conf = new(Config)
	err = config.Build().Get(conf)

	return
}

func tag(config *Config) *Tag {
	return &config.Tag
}
