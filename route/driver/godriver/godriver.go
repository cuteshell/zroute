package godriver

import (
	"zroute.io/model"
	"zroute.io/route/driver"
)

type GoDriver struct {
	model.Channel
	drivers map[string]driver.Driver
}

func New(c *model.Channel) *GoDriver {
	// godriver := driver.Driver{Channel: *c}
	// godriver.drivers["modbus-client"] = modbusc.New()
	// return &godriver
	return nil
}

func (d *GoDriver) OnCreate() error {
	return nil
}

func (d *GoDriver) OnScan() error {
	return nil
}

func (d *GoDriver) Command() error {
	return nil
}

func (d *GoDriver) SendFrame() error {
	return nil
}

func (d *GoDriver) RecvFrame() error {
	return nil
}

func (d *GoDriver) WaitResponse() error {
	return nil
}

func (d *GoDriver) OnDestroy() error {
	return nil
}
