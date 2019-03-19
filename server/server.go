package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/module"
	"github.com/qinhan-shu/gp-server/service/auth"
	"github.com/qinhan-shu/gp-server/service/config"
	"github.com/qinhan-shu/gp-server/service/gate/gin"
	"github.com/qinhan-shu/gp-server/service/manage"
)

var (
	port     = flag.String("addr", ":9999", "listen address")
	certFile = flag.String("certFile", "", "ssl certficate filename")
	keyFile  = flag.String("keyFile", "", "ssl private key filename")
	logLevel = flag.String("log-level", "error", "log level, optional( debug | info | warn | error | dpanic | panic | fatal), default is error")
)

func main() {
	flag.Parse()

	// init logger
	logger.InitLogger(*logLevel)

	gateService := gate.NewService(*port)
	if *certFile != "" && *keyFile != "" {
		gateService.AddTLSConfig(*certFile, *keyFile)
	}

	// register modules
	registerModule(gateService)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		logger.Sugar.Infof("Shutting down gate server...")
		gateService.Stop()
	}()

	logger.Sugar.Infof("Starting gate server on %s", *port)
	gateService.Start()
}

func registerModule(gate module.Gate) {
	c := config.NewConfig()
	dataStorage, err := c.GetDataStorage()
	if err != nil {
		logger.Sugar.Fatalf("failed to get data storage : %v", err)
	}

	auth.Register(gate, dataStorage)
	manage.Register(gate, dataStorage)
}
