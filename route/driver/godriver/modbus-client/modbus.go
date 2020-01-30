package modbusc

import (
	"github.com/sirupsen/logrus"
	godriver "zroute.io/route/driver/godriver"
)

type ModbusClient struct {
	godriver.GoDriver
}

func (m *ModbusClient) OnCreate() {
	channel := m.GoDriver.Channel
	if len(channel.RtuList) < 0 {
		logrus.Error()
	}
}
