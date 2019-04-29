package db

import (
	"github.com/qinhan-shu/gp-server/model/xorm"
)

// AddSubmitRecord : add new record of submit
func (m *MysqlDriver) AddSubmitRecord(submit *model.UserProblem) error {
	i, err := m.conn.Insert(submit)
	if err != nil {
		return err
	}
	if i == 0 {
		return ErrNoRowsAffected
	}
	return nil
}
