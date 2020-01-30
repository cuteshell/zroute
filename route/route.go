package route

import (
	"context"

	"github.com/coreos/etcd/store"
	"zroute.io/model"
	"zroute.io/route/channel"

	_ "zroute.io/utils/log"
)

type Route struct {
	ctx     context.Context
	running bool
	store   store.Store
	model   *model.Model
	chans   []*channel.Channel
	rtype   uint8
}

func New(ctx context.Context, project *Project) *Route {
	r := &Route{}
	r.model = model.New()

	var chans []model.Channel
	r.model.Preload("LinkList").Preload("RtuList").
		Preload("RtuList.PntList").Preload("RtuList.Parameter").Find(&chans)

	for _, c := range chans {
		channel := channel.New(ctx)
		channel.Channel = c
		r.chans = append(r.chans, channel)
	}
	return r
}

/*
	Ctx   context.Context
	Trans transport.Transport
	Rtus  map[int]rtu.Rtu // key is rtu_id
	Sendc chan []byte
*/
func (r *Route) Start() error {
	for _, c := range r.chans {
		c.Start()
	}
	return nil
}

func (r *Route) Close() error {
	r.model.Close()
	for _, c := range r.chans {
		c.Close()
	}
	return nil
}
