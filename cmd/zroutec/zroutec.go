package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
	"zroute.io/route"

	_ "zroute.io/utils/log"
)

var (
	version = flag.Bool("version", false, "Show current version.")
	join    = flag.Bool("join", false, "Run as join mode.")
)

func startZRouteC(ctx context.Context, project *route.Project) (*route.Route, error) {
	if project == nil {
		fmt.Println("haven't load project")
		return nil, errors.New("project is nil")
	}
	zroutec := route.New(ctx, project)
	err := zroutec.Start()
	return zroutec, err
}

func main() {
	flag.Parse()

	_, err := route.LoadConfig()
	if err != nil {
		log.Error("LoadConfig failed!")
		os.Exit(-1)
	}

	var project *route.Project
	if *join {
		log.Info("zroutec start as join mode")
	} else {
		project, err = route.LoadProject()
		if err != nil {
			log.Error("Load project failed")
			os.Exit(-1)
		}
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	zroutec, err := startZRouteC(ctx, project)
	if err != nil {
		log.Error("Failed to start", err)
		os.Exit(-1)
	}
	defer zroutec.Close()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, os.Kill, syscall.SIGTERM)
	log.Info("zroutec is runing...")
	<-signals
	log.Info("zroutec exit.")
}
