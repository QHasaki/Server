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
		resp.Code = protocol.Code_PERMISSION_DENIED
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
		resp.Code = protocol.Code_PERMISSION_DENIED
		return resp
	}

	token, err := a.cache.UpdateToken(user.ID)
	if err != nil {
		resp.Code = protocol.Code_INTERNAL
		return resp
	}

	resp.User = &protocol.UserInfo{
		Name:      user.Name,
		Sex:       user.Sex,
		Academy:   user.Academy,
		Major:     user.Major,
		Email:     user.Email,
		LastLogin: user.LastLogin,
		Role:      protocol.Role(user.Role),
	}
	resp.Token = token

	return resp
}

// Logout : log out, and del token
func (a *Auth) Logout(args map[string]interface{}) interface{} {
	resp := &protocol.LogOutResp{}
	if err := utils.CheckArgs(args, module.Request, module.Token); err != nil {
		resp.Code = protocol.Code_PERMISSION_DENIED
		return resp
	}

	req := &protocol.LogOutReq{}
	if err := proto.Unmarshal(parse.Bytes(args[module.Request]), req); err != nil {
		logger.Sugar.Errorf("failed to unmarshal : %v", err)
		resp.Code = protocol.Code_INVAILD_DATA
		return resp
	}

	// authority authentication
	userID, err := a.cache.GetUserIDByToken(parse.String(args[module.Token]))
	if err != nil {
		resp.Code = protocol.Code_PERMISSION_DENIED
		return resp
	}

	if err := a.cache.DelToken(userID); err != nil {
		resp.Code = protocol.Code_INTERNAL
	}
	return resp
}
