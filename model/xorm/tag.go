package model

type Tag struct {
	Id     int    `xorm:"not null pk autoincr INT(11)"`
	Detail string `xorm:"not null VARCHAR(100)"`
}
