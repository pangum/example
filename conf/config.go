package conf

import (
	`github.com/pangum/pangu`
)

// Config 配置外层包装
type Config struct {
	// 程序主体配置
	Agent Agent `json:"agent" yaml:"agent" xml:"agent" toml:"agent" validate:"required"`
}

func config(config *pangu.Config) (conf Config, err error) {
	err = config.Load(&conf)

	return
}

func agent(config Config) Agent {
	return config.Agent
}
