package channel

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"zroute.io/route/rtu"
	"zroute.io/route/transport"
	_ "zroute.io/utils/log"
)

// Channel contains rtu
type Channel struct {
	ID        string
	Name      string
	Address   string
	Ctx       context.Context
	transport transport.Transport
	Rtus      map[uint]*rtu.Rtu // key is rtu_id
	sendc     chan []byte
}

// New a channel
func New(ctx context.Context) *Channel {
	return &Channel{
		Ctx:   ctx,
		sendc: make(chan []byte, 100),
	}
}

// Start the channel
func (c *Channel) Start() (err error) {
	err = c.transport.Open()
	if err != nil {
		log.Error("Open trans failed")
		return
	}
	for id, r := range c.Rtus {
		c.Rtus[id] = rtu.New(r, c.sendc)
	}
	defer func(rtus map[uint]*rtu.Rtu) {
		for _, rtu := range rtus {
			rtu.OnDestroy()
		}
	}(c.Rtus)
	go func(c *Channel) {
		for {
			select {
			case <-c.Ctx.Done():
				log.Info("channel routine was exit")
				return
			case data, ok := <-c.sendc:
				if !ok {
					log.Info("channel has been closed")
				}
				n, err := c.transport.Write(data)
				if err != nil {
					log.Error("send error")
					continue
				}
				log.Debugf("send %d byte data", n)
			default:
				for _, rtu := range c.Rtus {
					rtu.OnScan()
				}
				time.Sleep(10 * time.Millisecond)
			}
		}
	}(c)
	return
}

// Close the channel
func (c *Channel) Close() (err error) {
	err = c.transport.Close()
	if err != nil {
		log.Error("Close trans failed")
		return
	}
	return
}
