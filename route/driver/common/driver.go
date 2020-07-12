package common

import (
	"log"
	"net/rpc"
	"os/exec"

	"github.com/hashicorp/go-plugin"
	"zroute.io/model"
	"zroute.io/route/driver/common"
)

// Driver protocol to communicate
// type Driver interface {
// 	OnCreate() error
// 	OnScan() error
// 	Command() error
// 	SendFrame() error
// 	RecvFrame() error
// 	WaitResponse() error
// 	OnDestroy() error
// }

// type DriverManager struct {
// }

type Driver interface {
	ReadPoints(req Request) (Response, error)
}

type Request struct {
	PntList []model.Pnt
}

type Response struct {
	Data []byte
}

type DriverPlugin struct {
	Impl Driver
}

func (d *DriverPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &DriverServer{Impl: d.Impl}, nil
}

func (*DriverPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &DriverClient{client: c}, nil
}

func LoadDriver(driverName string) (*plugin.Client, Driver) {
	driverPath := "../../route/driver/module/" + driverName + ".mo"
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: common.Handshake,
		Plugins: map[string]plugin.Plugin{
			driverName: &common.DriverPlugin{},
		},
		Cmd: exec.Command(driverPath),
	})
	defer client.Kill()
	rpcClient, err := client.Client()
	if err != nil {
		log.Fatal(err)
	}
	raw, err := rpcClient.Dispense(driverName)
	if err != nil {
		log.Fatal(err)
	}
	driver := raw.(common.Driver)
	return client, driver
}
