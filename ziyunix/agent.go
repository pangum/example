package ziyunix

import (
	`time`

	`agent/ziyunix/dispatcher`

	`github.com/pangum/logging`
	`github.com/pangum/mqtt`
	`github.com/pangum/pangu`
)

type (
	// Agent 执行器
	Agent struct {
		convert *dispatcher.Convert

		mqtt   *mqtt.Client
		logger *logging.Logger
	}

	agentIn struct {
		pangu.In

		Convert *dispatcher.Convert
		Mqtt    *mqtt.Client
		Logger  *logging.Logger
	}
)

func newAgent(in agentIn) *Agent {
	return &Agent{
		convert: in.Convert,
		mqtt:    in.Mqtt,
		logger:  in.Logger,
	}
}

func (a *Agent) Start() (err error) {
	if err = a.mqtt.Subscribe(`convert/+`, a.convert, mqtt.Qos2()); nil != err {
		return
	}

	return
}

func (a *Agent) Stop() error {
	return a.mqtt.Disconnect(time.Second)
}

func (a *Agent) Name() string {
	return `子云执行器`
}
