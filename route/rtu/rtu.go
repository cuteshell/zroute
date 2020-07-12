package rtu

import (
	"github.com/hashicorp/go-plugin"
	log "github.com/sirupsen/logrus"
	"zroute.io/model"
	"zroute.io/route/driver/common"
)

// Rtu represet a device
type Rtu struct {
	model.Rtu
	Client *plugin.Client
	Driver common.Driver
	sendc  chan []byte
}

// New create a new rtu
func New(r *model.Rtu, sendc chan []byte) *Rtu {
	rtu := Rtu{Rtu: *r, sendc: sendc}
	rtu.Client, rtu.Driver = common.LoadDriver("modbus-tcp")
	rtu.OnCreate()
	log.Debug("modbus-tcp RTU was created")
	return &rtu
}

// OnScan :query the data
func (r *Rtu) OnScan() error {
	req := common.Request{PntList: r.PntList}
	data, err := r.Driver.ReadPoints(req)
	if err != nil {
		r.sendc <- data.Data
	}
	return err
}

// when rtu was created
func (r *Rtu) OnCreate() error {
	log.Debug("OnCreate was called")
	return nil
}

// when rtu was descroyed
func (r *Rtu) OnDestroy() error {
	log.Debug("OnDestry was called")
	r.Client.Kill()
	return nil
}
