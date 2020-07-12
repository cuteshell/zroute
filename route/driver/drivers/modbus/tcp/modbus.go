package main

import (
	"fmt"

	"github.com/hashicorp/go-plugin"
	"zroute.io/route/driver/common"
)

type ModbusTCP struct {
}

func (m *ModbusTCP) ReadPoints(req common.Request) (common.Response, error) {
	var resp common.Response
	fmt.Println("ReadPoints was called")
	return resp, nil
}

func main() {
	driver := &ModbusTCP{}
	driverName := "modbus-tcp"

	var pluginMap = map[string]plugin.Plugin{
		driverName: &common.DriverPlugin{Impl: driver},
	}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: common.Handshake,
		Plugins:         pluginMap,
	})
}
