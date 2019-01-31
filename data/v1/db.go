package data

import (
	"github.com/QHasaki/Server/data/v1/driver/mysql"
	"github.com/QHasaki/Server/logger"
)

type DB struct {
	origin *originDB
	cache  *Cache
}

func (p *DB) Set(document string, data Data, where Data) error {
	if err := getCondition(document, data, where); err != nil {
		return err
	}
	cacheKey := p.MakeCacheKey(document, where)
	if err := driver.Set(p.origin.conn, document, data, where); err != nil {
		return err
	}
	if err := p.cache.Delete(cacheKey); err != nil {
		logger.Sugar.Errorf("cant delete data cache : %v", err)
	}
	return nil
}

func (p *DB) Get(document string, column []string, where Data) (Data, error) {
	if err := setCondition(document, column, where); err != nil {
		return nil, err
	}
	cacheKey := p.MakeCacheKey(document, where)
	cacheData, err := p.cache.GetAll(cacheKey)
	if err != nil || len(cacheData) < 1 {
		data, err := driver.Get(p.origin.conn, document, []string{"*"}, where)
		if err != nil {
			return nil, err
		}
		err = p.cache.SetAll(cacheKey, data)
		if err != nil {
			logger.Sugar.Errorf("cant save to data cache: %v", err)
		}
		return data, nil
	}
	return cacheData, nil
}

func (p *DB) Inc(document string, column []string, where Data) error {
	if err := driver.Inc(p.origin.conn, document, column, where); err != nil {
		return err
	}
	return nil
}
