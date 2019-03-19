package auth

import (
	"net/http"

	"github.com/golang/protobuf/proto"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/protocol"
	"github.com/qinhan-shu/gp-server/utils"
)

// Login : authentication, and get token
func (a *Auth) Login(r *http.Request) (int, proto.Message) {
	// get request and response
	code := http.StatusOK
	req := &protocol.LoginReq{}
	resp := &protocol.LoginResp{}
	data, err := utils.GetRequestBody(r)
	if err != nil {
		code = http.StatusBadRequest
		return code, resp
	}

	if err := proto.Unmarshal(data, req); err != nil {
		logger.Sugar.Errorf("failed to unmarshal : %v", err)
		code = http.StatusBadRequest
		return code, resp
	}

	user, err := a.db.CheckPlayer(req.Account, utils.MD5(req.Password))
	if err != nil {
		code = http.StatusUnauthorized
		return code, resp
	}

	token, err := a.cache.UpdateToken(user.ID)
	if err != nil {
		code = http.StatusInternalServerError
		return code, resp
	}

	resp.User = user.TurnProto()
	resp.Token = token

	return code, resp
}

// Logout : log out, and del token
func (a *Auth) Logout(r *http.Request) (int, proto.Message) {
	code := http.StatusOK
	token, err := utils.GetToken(r)
	if err != nil {
		code = http.StatusBadRequest
		return code, nil
	}

	a.cache.DelTokenByToken(token) // nolint : err check
	return code, nil
}
