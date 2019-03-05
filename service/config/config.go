package config

import (
	"fmt"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/module"
	"github.com/qinhan-shu/gp-server/service/cache/redis"
	"github.com/qinhan-shu/gp-server/service/db/gorm/mysql"
	"github.com/qinhan-shu/gp-server/utils/parse"
)

// InitConfig is to get config
// Fatal if failed to get config
func (c *Config) InitConfig() {
	configMap, err := c.configSource.GetConfig()
	if err != nil {
		logger.Sugar.Fatalf("failed to init config : %v", err)
	}

	c.Lock()
	defer c.Unlock()

	c.configMap = configMap
}

// ReloadConfig is to reload config
// error if failed to get config
func (c *Config) ReloadConfig() error {
	configMap, err := c.configSource.GetConfig()
	if err != nil {
		return err
	}

	c.Lock()
	defer c.Unlock()

	c.configMap = configMap

	return nil
}

// GetConfigByName is to get config value by config key
func (c *Config) GetConfigByName(configName string) (string, error) {
	c.RLock()
	defer c.RUnlock()

	configValue := c.configMap[configName]
	if configValue == "" {
		return "", fmt.Errorf("missing config : %s", configName)
	}

	return configValue, nil
}

// GetDataStorage : new data storage system
func (c *Config) GetDataStorage() (*module.DataStorage, error) {
	keys := []string{
		"REDIS_ADDR",
		"REDIS_PASS",
		"REDIS_POOLSIZE",
		"MYSQL_ADDR",
		"MYSQL_USER",
		"MYSQL_PASS",
		"MYSQL_DBNAME",
		"MYSQL_OPEN_CONNS_NUM",
		"MYSQL_IDLE_CONNS_NUM",
	}
	values := make(map[string]string)

	for _, key := range keys {
		value, err := c.GetConfigByName(key)
		if err != nil {
			return nil, err
		}
		values[key] = value
	}

	redisPoolSize, err := parse.IntWithError(values["REDIS_POOLSIZE"])
	if err != nil {
		return nil, err
	}

	openConnsNum, err := parse.IntWithError(values["MYSQL_OPEN_CONNS_NUM"])
	if err != nil {
		return nil, err
	}

	idleConnsNum, err := parse.IntWithError(values["MYSQL_IDLE_CONNS_NUM"])
	if err != nil {
		return nil, err
	}

	mysqlConfig := &db.MysqlConfig{
		Addr:            values["MYSQL_ADDR"],
		Username:        values["MYSQL_USER"],
		Password:        values["MYSQL_PASS"],
		DBName:          values["MYSQL_DBNAME"],
		MaxOpenConnsNum: int(openConnsNum),
		MaxIdleConnsNum: int(idleConnsNum),
	}
	mysqlDB, err := db.NewMysqlDriver(mysqlConfig)
	if err != nil {
		return nil, err
	}

	redisConfig := &cache.RedisConfig{
		Addr:     values["REDIS_ADDR"],
		Password: values["REDIS_PASS"],
		PoolSize: int(redisPoolSize),
	}
	redisCache, err := cache.NewRedisCache(redisConfig)
	if err != nil {
		return nil, err
	}
	return &module.DataStorage{
		DB:    mysqlDB,
		Cache: redisCache,
	}, nil
}
