package config

import (
	"sync"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/module"
	github "github.com/qinhan-shu/gp-server/service/config/source/github"
	local "github.com/qinhan-shu/gp-server/service/config/source/local"
)

// Config describes Config model
type Config struct {
	rwMutex      sync.RWMutex
	configSource module.ConfigSource
	configMap    module.ConfigMap
}

// NewConfig is the constructor of config model
func NewConfig() module.Config {
	var (
		source module.ConfigSource
		err    error
	)
	logger.Sugar.Info("get config from source (github)")
	source, err = github.NewConfigSource()
	if err != nil {
		logger.Sugar.Errorf("failed to get config from source (github) : %v", err)
		logger.Sugar.Info("get config from source (local file)")
		source, err = local.NewConfigSource()
		if err != nil {
			logger.Sugar.Errorf("failed to get config from source (local file) : %v", err)
			logger.Sugar.Fatalf("failed to get config from source(github & local file)")
		}
	}

	c := &Config{
		configSource: source,
	}

	c.InitConfig()

	return c
}
