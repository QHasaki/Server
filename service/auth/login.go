package auth

import (
	"github.com/golang/protobuf/proto"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/module"
	"github.com/qinhan-shu/gp-server/protocol"
	"github.com/qinhan-shu/gp-server/utils"
	"github.com/qinhan-shu/gp-server/utils/parse"
)

// Login : authentication, and get token
func (a *Auth) Login(args map[string]interface{}) interface{} {
	resp := &protocol.LoginResp{}
	if err := utils.CheckArgs(args, module.Request); err != nil {
		resp.Code = protocol.Code_INVAILD_DATA
		return resp
	}

	req := &protocol.LoginReq{}
	if err := proto.Unmarshal(parse.Bytes(args[module.Request]), req); err != nil {
		logger.Sugar.Errorf("failed to unmarshal : %v", err)
		resp.Code = protocol.Code_INVAILD_DATA
		return resp
	}

	user, err := a.db.GetUserByAuthCode(req.Username)
	if err != nil {
		resp.Code = protocol.Code_INVAILD_DATA
		return resp
	}

	token, err := a.cache.UpdateToken(user.ID)
	if err != nil {
		resp.Code = protocol.Code_INTERNAL
		return resp
	}

	resp.Token = token
	return resp
}

// Logout : log out, and del token
func (a *Auth) Logout(args map[string]interface{}) interface{} {
	return nil
}
