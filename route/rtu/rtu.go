package rtu

import (
	"github.com/hashicorp/go-plugin"
	log "github.com/sirupsen/logrus"
	"zroute.io/route/driver/common"
	"zroute.io/route/pnt"
)

// Rtu represet a device
type Rtu struct {
	ID     string
	Name   string
	Client *plugin.Client
	Driver common.Driver
	sendc  chan []byte
	Pnts   map[string]pnt.Pnt
}

// New create a new rtu
func New(r *Rtu, sendc chan []byte) *Rtu {
	rtu := Rtu{ID: r.ID, Name: r.Name, sendc: sendc}
	rtu.Client, rtu.Driver = common.LoadDriver("modbus-tcp")
	rtu.OnCreate()
	log.Debug("modbus-tcp RTU was created")
	return &rtu
}

// OnScan :query the data
func (r *Rtu) OnScan() error {
	req := common.Request{Pnts: r.Pnts}
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
