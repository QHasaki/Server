package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//"github.com/qinhan-shu/gp-server/model/gorm"
)

// MysqlConfig : mysql 数据库配置
type MysqlConfig struct {
	MaxConnection int // 最大连接数
	Addr          string
	Username      string
	Password      string
	DBName        string
}

func (c *MysqlConfig) getDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true",
		c.Username, c.Password, c.Addr, c.DBName)
}

// MysqlDriver : mysql 驱动
type MysqlDriver struct {
	config *MysqlConfig
	conn   *gorm.DB
}

// NewMysqlDriver : MysqlDriver 的构造方法
func NewMysqlDriver(config *MysqlConfig) (*MysqlDriver, error) {
	mysqlDriver := &MysqlDriver{
		config: config,
	}

	if err := mysqlDriver.Connect(); err != nil {
		return nil, err
	}

	return mysqlDriver, nil
}

// Connect : 连接数据库
func (m *MysqlDriver) Connect() error {
	if m.config == nil {
		return ErrNoMysqlConf
	}

	db, err := gorm.Open("mysql", m.config.getDSN())
	if err != nil {
		return err
	}

	// 设置表名就是结构体的名字
	// 如果不设置的话，表名默认为结构体名的复数
	db.SingularTable(true)

	m.conn = db
	m.autoMigrate()

	return nil
}

// 自动迁移，对于注册的model， 将增加数据库中没有但是model中定义过的字段，不会删除（改变）原先的字段&数据
func (m *MysqlDriver) autoMigrate() {
	//m.conn.AutoMigrate(&model.Test{})
}