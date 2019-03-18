package config

import (
	"sync"

	"github.com/qinhan-shu/gp-server/module"
	"github.com/qinhan-shu/gp-server/service/config/source/github"
)

// Config describes Config model
type Config struct {
	rwMutex      sync.RWMutex
	configSource module.ConfigSource
	configMap    module.ConfigMap
}

// NewConfig is the constructor of config model
func NewConfig() module.Config {
	c := &Config{
		configSource: source.NewGithub(),
	}

	c.InitConfig()

	return c
}
