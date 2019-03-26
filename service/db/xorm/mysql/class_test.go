package db_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	model_utils "github.com/qinhan-shu/gp-server/model"
	"github.com/qinhan-shu/gp-server/model/xorm"
)

func TestMysqlDriver_GetClassesNum(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	num, err := mysqlDriver.GetClassNum()
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("the num of class : %d", num)
}

func TestMysqlDriver_AddClass(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	class := &model.Class{
		CreateTime:   time.Now().Unix(),
		Introduction: "这个一个测试的班级",
		IsCheck:      0,
		Name:         "测试班级",
		Tutor:        1,
	}
	announcements := []*model.Announcement{
		&model.Announcement{
			CreateTime: time.Now().Unix(),
			Detail:     "大家进入班级请改名",
			Publisher:  1,
		},
		&model.Announcement{
			CreateTime: time.Now().Unix(),
			Detail:     "大家进入班级请改名111",
			Publisher:  1,
		},
	}

	if err := mysqlDriver.AddClass(&model_utils.IntactClass{
		Class:         class,
		Announcements: announcements,
	}); err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v\n", class)
}

func TestMysqlDriver_GetClasses(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	var pageIndex int64 = 1
	var pageNum int64 = 3
	classes, err := mysqlDriver.GetClasses(pageNum, pageIndex)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(len(classes))
	for _, class := range classes {
		t.Logf("%+v\n", class.Class)
		t.Logf("%+v\n", class.Announcements)
	}
}

func TestMysqlDriver_GetClassByID(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	var classID int64 = 1
	intactClass, err := mysqlDriver.GetClassByID(classID)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v\n", intactClass.Class)
	t.Logf("%+v\n", intactClass.Announcements)

}

func TestMysqlDriver_UpdateClass(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	var classID int64 = 1
	originClass, err := mysqlDriver.GetClassByID(classID)
	if err != nil {
		t.Error(err)
		return
	}

	change := &model_utils.IntactClass{
		Class: &model.Class{
			Id:   classID,
			Name: originClass.Class.Name + "000",
		},
		Announcements: []*model.Announcement{
			&model.Announcement{
				CreateTime: time.Now().Unix(),
				Detail:     "这个公告是在测试中增加的",
				Publisher:  1,
			},
		},
	}

	if err := mysqlDriver.UpdateClass(change); err != nil {
		t.Error(err)
		return
	}

	changedClass, err := mysqlDriver.GetClassByID(classID)
	if err != nil {
		t.Error(err)
		return
	}

	if !assert.NotEqual(t, originClass.Class.Name, changedClass.Class.Name) {
		t.Error("filed [Name] is not changed")
		return
	}

	if !assert.Equal(t, changedClass.Class.Name, change.Class.Name) {
		t.Error("filed [Name] is not changed to expected value")
		return
	}

	if !assert.Equal(t, len(changedClass.Announcements), len(originClass.Announcements)+1) {
		t.Error("filed [Announcements] is not changed to expected value")
		return
	}

	if !assert.Equal(t, originClass.Class.Introduction, changedClass.Class.Introduction) {
		t.Error("other filed [Introduction] is changed")
		return
	}
}

func TestMysqlDriver_MemberManage(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	var manageType int64 = 1
	var classID int64 = 1
	var memberID int64 = 2
	if err := mysqlDriver.MemberManage(manageType, classID, memberID); err != nil {
		t.Error(err)
		return
	}
}

func TestAddSomeClasses(t *testing.T) {
	num := 10
	for i := 0; i < num; i++ {
		TestMysqlDriver_AddClass(t)
	}
}
