package user_manage_test

import (
	"testing"

	"github.com/golang/protobuf/proto"

	"github.com/qinhan-shu/gp-server/module"
	"github.com/qinhan-shu/gp-server/protocol"
	"github.com/qinhan-shu/gp-server/service/config"
	"github.com/qinhan-shu/gp-server/service/user_manage"
	"github.com/qinhan-shu/gp-server/utils/mode"
)

func TestUserManage_GetUsers(t *testing.T) {
	mode.SetMode(mode.TestMode)
	dataStorage, err := config.NewConfig().GetDataStorage()
	if err != nil {
		t.Error(err)
		return
	}
	userManagerModule := user_manage.NewUserManager(dataStorage)

	// get user operations
	reqBytes, err := proto.Marshal(&protocol.GetUsersReq{
		GetAll: false,
		Role:   protocol.Role_TEACHER,
	})
	if err != nil {
		t.Error(err)
		return
	}
	args := map[string]interface{}{
		module.Token:   "1",
		module.Request: reqBytes,
	}

	data := userManagerModule.GetUsers(args)
	resp := data.(*protocol.GetUsersResp)
	if resp.Code != protocol.Code_OK {
		t.Error(err)
		return
	}
	for _, user := range resp.Users {
		t.Logf("%+v", user)
	}
}

func TestUserManage_AddUsers(t *testing.T) {
	mode.SetMode(mode.TestMode)
	dataStorage, err := config.NewConfig().GetDataStorage()
	if err != nil {
		t.Error(err)
		return
	}
	userManagerModule := user_manage.NewUserManager(dataStorage)

	// add users operations
	reqBytes, err := proto.Marshal(&protocol.AddUsersReq{
		Users: []*protocol.UserInfo{
			&protocol.UserInfo{
				Name:     "asdf",
				Sex:      true,
				Role:     0,
				Academy:  "it",
				Major:    "cs",
				Username: "AddUsers_10",
				Password: "AddUsers_10",
			},
			&protocol.UserInfo{
				Name:     "asdf",
				Sex:      true,
				Role:     0,
				Academy:  "it",
				Major:    "cs",
				Username: "AddUsers_20",
				Password: "AddUsers_20",
			},
			&protocol.UserInfo{
				Name:     "asdf",
				Sex:      true,
				Role:     0,
				Academy:  "it",
				Major:    "cs",
				Username: "AddUsers_30",
				Password: "AddUsers_30",
			},
			&protocol.UserInfo{
				Id:       2,
				Name:     "asdf",
				Sex:      false,
				Role:     0,
				Academy:  "it",
				Major:    "cs",
				Username: "AddUsers_40",
				Password: "AddUsers_40",
			},
		},
	})
	if err != nil {
		t.Error(err)
		return
	}
	args := map[string]interface{}{
		module.Token:   "1",
		module.Request: reqBytes,
	}

	data := userManagerModule.AddUsers(args)
	resp := data.(*protocol.AddUsersResp)
	if resp.Code != protocol.Code_OK {
		t.Error(err)
		return
	}

	t.Log("-------Succeed-------")
	for _, user := range resp.Succeed {
		t.Logf("%+v", user)
	}

	t.Log("-------failed-------")
	for _, user := range resp.Fail {
		t.Logf("%+v", user)
	}
}

func TestUserManage_UpdateUsers(t *testing.T) {
	mode.SetMode(mode.TestMode)
	dataStorage, err := config.NewConfig().GetDataStorage()
	if err != nil {
		t.Error(err)
		return
	}
	userManagerModule := user_manage.NewUserManager(dataStorage)

	// update users operations
	reqBytes, err := proto.Marshal(&protocol.UpdateUsersReq{
		Users: []*protocol.UserInfo{
			&protocol.UserInfo{
				Id:       2,
				Name:     "1",
				Sex:      true,
				Role:     0,
				Academy:  "2",
				Major:    "2",
				Username: "2",
				Password: "2",
			},
		},
	})
	if err != nil {
		t.Error(err)
		return
	}
	args := map[string]interface{}{
		module.Token:   "1",
		module.Request: reqBytes,
	}

	data := userManagerModule.UpdateUsers(args)
	resp := data.(*protocol.UpdateUsersResp)
	if resp.Code != protocol.Code_OK {
		t.Error(err)
		return
	}

	t.Log("-------Succeed-------")
	for _, user := range resp.Succeed {
		t.Logf("%+v", user)
	}

	t.Log("-------failed-------")
	for _, user := range resp.Fail {
		t.Logf("%+v", user)
	}
}

func TestUserManage_DelUsers(t *testing.T) {
	mode.SetMode(mode.TestMode)
	dataStorage, err := config.NewConfig().GetDataStorage()
	if err != nil {
		t.Error(err)
		return
	}
	userManagerModule := user_manage.NewUserManager(dataStorage)

	// del users operations
	reqBytes, err := proto.Marshal(&protocol.DelUsersReq{
		Users: []*protocol.UserInfo{
			&protocol.UserInfo{
				Id: 14,
			},
			&protocol.UserInfo{
				Id: 15,
			},
		},
	})
	if err != nil {
		t.Error(err)
		return
	}
	args := map[string]interface{}{
		module.Token:   "1",
		module.Request: reqBytes,
	}

	data := userManagerModule.DelUsers(args)
	resp := data.(*protocol.DelUsersResp)
	if resp.Code != protocol.Code_OK {
		t.Error(err)
		return
	}

	t.Log("-------Succeed-------")
	for _, user := range resp.Succeed {
		t.Logf("%+v", user)
	}

	t.Log("-------failed-------")
	for _, user := range resp.Fail {
		t.Logf("%+v", user)
	}
}
