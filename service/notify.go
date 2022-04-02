package service

import (
	"github.com/goexl/gox"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/pangum/mqtt"
	"github.com/ziyunix/core"
	"github.com/ziyunix/message/task"
)

// Notify 通知
type Notify struct {
	mqtt *mqtt.Client

	_ gox.CannotCopy
}

func newNotify(mqtt *mqtt.Client) *Notify {
	return &Notify{
		mqtt: mqtt,
	}
}

func (n *Notify) Notification(id int64, status core.Status, timestamp *timestamp.Timestamp) {
	notification := &task.Convert{
		Id:        id,
		Status:    status,
		Timestamp: timestamp,
	}

	if err := n.mqtt.Publish(``, notification); nil != err {
		panic(err)
	}
}
