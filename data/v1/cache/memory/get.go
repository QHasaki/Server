package cache

import (
	"github.com/QHasaki/Server/data/v1/error"
)

func (c *Cache) GetAll(key string) (map[string]interface{}, error) {
	data, ok := c.storage.Load(key)
	if ok && data != nil {
		switch data.(type) {
		case *CacheData:
			return data.(*CacheData).Data, nil
		case nil:
			return make(map[string]interface{}), nil
		default:
			return nil, dataError.ErrInvalidType
		}
	} else {
		return nil, dataError.ErrNoRowsFound
	}
}
