package store

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"go.etcd.io/etcd/clientv3"
)

var (
	dialTimeout    = 5 * time.Second
	requestTimeout = 10 * time.Second
	endpoints      = []string{"localhost:2379", "localhost:22379", "localhost:32379"}
)

type Store struct {
	cli   clientv3.Client
	ctx   context.Context
}

func New() *Store {
	s := &Store{}
	s.cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: dialTimeout,
	})
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer s.cli.Close()
	s.ctx = context.TODO()
	return s
}

func (s *Store) Put(key, val string) (interface{}, error){
	return cli.Put(s.ctx, key, val)
}

func (s *Store) Get(key, val string) (interface{}, error){
	return cli.Get(s.ctx, key, val)
}

func (s *Store) Delete(key string) (interface{}, error){
	return cli.Delete(s.ctx, key, val)
}

func (s *Store) Watch(key string) interface{}{
	return cli.Watch(s.ctx, key)
}