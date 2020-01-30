package route

import (
	"zroute.io/route/channel"
	"zroute.io/route/point"
	"zroute.io/route/rtu"
	"zroute.io/route/transport"
)

type Config struct {
}

type Project struct {
	Transport []transport.Transport
	Channel   []channel.Channel
	Device    []rtu.Rtu
	Point     []point.Point
}

func LoadConfig() (*Config, error) {
	config := new(Config)

	return config, nil
}

func LoadProject() (*Project, error) {
	project := new(Project)

	return project, nil
}
