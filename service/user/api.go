package user

import (
	"net/http"
	"time"

	"github.com/golang/protobuf/proto"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/model/transform"
	"github.com/qinhan-shu/gp-server/model/xorm"
	"github.com/qinhan-shu/gp-server/protocol"
	"github.com/qinhan-shu/gp-server/utils"
)

// GetUsers : get users
func (u *User) GetUsers(r *http.Request) proto.Message {
	req := &protocol.GetUsersReq{}
	resp := &protocol.GetUsersResp{Status: &protocol.Status{}}

	_, status := u.checkArgsandAuth(r, req)
	if status != nil {
		logger.Sugar.Error(status.Message)
		resp.Status = status
		return resp
	}

	var users []*model.User
	var err error
	// get required users informations
	role := req.Role
	if req.Role == 0 {
		users, err = u.db.GetUsers(req.PageNum, req.PageIndex)
	} else {
		users, err = u.db.GetUsersByRole(req.PageNum, req.PageIndex, role)
	}
	if err != nil {
		logger.Sugar.Errorf("failed to get users : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "failed to get users"
		return resp
	}

	// add users informations to response
	for _, user := range users {
		resp.Users = append(resp.Users, transform.UserToProto(user))
	}

	// get all number
	usersNum, err := u.db.GetUsersNum(req.Role)
	if err != nil {
		logger.Sugar.Errorf("failed to get the number of users : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "failed to get the number of users"
		return resp
	}
	resp.Total = usersNum
	resp.PageIndex = req.PageIndex
	resp.PageNum = req.PageNum
	return resp
}

// AddUsers : add users to db
func (u *User) AddUsers(r *http.Request) proto.Message {
	req := &protocol.AddUsersReq{}
	resp := &protocol.AddUsersResp{Status: &protocol.Status{}}

	_, status := u.checkArgsandAuth(r, req)
	if status != nil {
		logger.Sugar.Error(status.Message)
		resp.Status = status
		return resp
	}

	// add users
	ts := time.Now().Unix()
	for _, protoUser := range req.Users {
		user := transform.ProtoToUser(protoUser)
		user.LastLogin = ts
		user.Create = ts
		if err := u.db.AddUser(user); err != nil {
			resp.Fail = append(resp.Fail, protoUser)
		} else {
			resp.Succeed = append(resp.Succeed, transform.UserToProto(user))
		}
	}

	return resp
}

// UpdateUsers : update users
func (u *User) UpdateUsers(r *http.Request) proto.Message {
	req := &protocol.UpdateUsersReq{}
	resp := &protocol.UpdateUsersResp{Status: &protocol.Status{}}

	_, status := u.checkArgsandAuth(r, req)
	if status != nil {
		logger.Sugar.Error(status.Message)
		resp.Status = status
		return resp
	}

	// update users
	for _, protoUser := range req.Users {
		user := transform.ProtoToUser(protoUser)
		if err := u.db.UpdateUser(user); err != nil {
			resp.Fail = append(resp.Fail, protoUser)
		} else {
			resp.Succeed = append(resp.Succeed, transform.UserToProto(user))
		}
	}

	return resp
}

// DelUsers : delete users
func (u *User) DelUsers(r *http.Request) proto.Message {
	req := &protocol.DelUsersReq{}
	resp := &protocol.DelUsersResp{Status: &protocol.Status{}}

	_, status := u.checkArgsandAuth(r, req)
	if status != nil {
		logger.Sugar.Error(status.Message)
		resp.Status = status
		return resp
	}

	// delete users
	for _, userID := range req.UsersId {
		if err := u.db.DelUser(userID); err != nil {
			resp.Fail = append(resp.Fail, userID)
		} else {
			resp.Succeed = append(resp.Succeed, userID)
		}
	}

	return resp
}

// GetSubmitRecord : submit record
func (u *User) GetSubmitRecord(r *http.Request) proto.Message {
	req := &protocol.GetSubmitRecordReq{}
	resp := &protocol.GetSubmitRecordResp{Status: &protocol.Status{}}
	// get token and data
	data, token, err := utils.GetReqAndToken(r)
	if err != nil {
		logger.Sugar.Error(err)
		resp.Status.Code = protocol.Code_DATA_LOSE
		resp.Status.Message = err.Error()
		return resp
	}
	if err := proto.Unmarshal(data, req); err != nil {
		logger.Sugar.Errorf("failed to unmarshal request body : %v", err)
		resp.Status.Code = protocol.Code_DATA_LOSE
		resp.Status.Message = "failed to unmarshal request body"
		return resp
	}

	// check token
	userID, err := u.cache.GetUserIDByToken(token)
	if err != nil {
		logger.Sugar.Errorf("failed to get token : %v", err)
		resp.Status.Code = protocol.Code_UNAUTHORIZATED
		resp.Status.Message = "invalid token"
		return resp
	}

	records, num, err := u.db.GetSubmitRecord(userID, req.PageNum, req.PageIndex)
	if err != nil {
		logger.Sugar.Errorf("failed to get submit records : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "failed to get submit records"
		return resp
	}

	for _, v := range records {
		resp.SubmitRecords = append(resp.SubmitRecords, transform.UserProblemToProto(v))
	}
	resp.PageNum = req.PageNum
	resp.PageIndex = req.PageIndex
	resp.Total = num
	return resp
}
