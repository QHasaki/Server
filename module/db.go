package module

import (
	"github.com/qinhan-shu/gp-server/model/gorm"
)

// DB : 数据库接口
type DB interface {
	GetPlayerByID(id int) (model.Test, error)
}
