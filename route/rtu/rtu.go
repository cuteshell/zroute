package rtu

import (
	"zroute.io/model"
	"zroute.io/route/driver"
	godriver "zroute.io/route/driver/godriver"
)

// Rtu represet a device
type Rtu struct {
	model.Rtu
	Driver driver.Driver
}

// New create a new rtu
func New(r *model.Rtu) *Rtu {
	rtu := Rtu{Rtu: *r}
	rtu.Driver = godriver.New(&rtu.Channel)
	rtu.OnCreate()
	return &rtu
}

// OnScan :query the data
func (r *Rtu) OnScan() error {
	r.Driver.OnScan()
	return nil
}

func (r *Rtu) OnCreate() error {
	r.Driver.OnCreate()
	return nil
}

func (r *Rtu) OnDestroy() error {
	r.Driver.OnDestroy()
	return nil
}
