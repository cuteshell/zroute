package common

import (
	"log"
	"net/rpc"
	"os/exec"
	"zroute.io/route/pnt"

	"github.com/hashicorp/go-plugin"
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
	Pnts  map[string]pnt.Pnt
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
		HandshakeConfig: Handshake,
		Plugins: map[string]plugin.Plugin{
			driverName: &DriverPlugin{},
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
	driver := raw.(Driver)
	return client, driver
}
