package channel

import (
	"context"
	"fmt"
	"time"

	"zroute.io/model"
	"zroute.io/route/rtu"
	"zroute.io/route/transport"
)

// Channel contains rtu
type Channel struct {
	model.Channel
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
		fmt.Println("Open trans failed")
		return
	}
	for _, r := range c.RtuList {
		c.Rtus[c.ID] = rtu.New(&r)
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
				fmt.Println("channel routine was exit")
				return
			case data, ok := <-c.sendc:
				if !ok {
					fmt.Println("channel has been closed")
				}
				n, err := c.transport.Write(data)
				if err != nil {
					fmt.Println("send error")
					continue
				}
				fmt.Printf("send %d byte data", n)
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
		fmt.Println("Open trans failed")
		return
	}
	return
}
