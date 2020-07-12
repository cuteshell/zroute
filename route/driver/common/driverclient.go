package common

import "net/rpc"

type DriverClient struct {
	client *rpc.Client
}

func (d *DriverClient) ReadPoints(req Request) (Response, error) {
	var rsp Response
	err := d.client.Call("Plugin.ReadPoints", req, &rsp)
	return rsp, err
}
