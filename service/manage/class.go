package manage

import (
	"fmt"
	"net/http"

	"github.com/golang/protobuf/proto"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/protocol"
)

// GetClasses : get all classes
func (m *BackStageManage) GetClasses(r *http.Request) proto.Message {
	req := &protocol.GetClassesReq{}
	resp := &protocol.GetClassesResp{Status: &protocol.Status{}}

	// get token and data
	data, token, err := getReqAndToken(r)
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
	_, err = m.cache.GetUserIDByToken(token)
	if err != nil {
		logger.Sugar.Errorf("failed to get token : %v", err)
		resp.Status.Code = protocol.Code_UNAUTHORIZATED
		resp.Status.Message = "invalid token"
		return resp
	}

	classes, err := m.db.GetClasses(req.PageNum, req.PageIndex)
	if err != nil {
		logger.Sugar.Errorf("failed to get classes : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "failed to get classes"
		return resp
	}

	for _, class := range classes {
		resp.Classes = append(resp.Classes, class.TurnMinProto())
	}

	// get all number
	classesNum, err := m.db.GetClassNum()
	if err != nil {
		logger.Sugar.Errorf("failed to get the number of classes : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "failed to get the number of classes"
		return resp
	}

	resp.Total = classesNum
	resp.PageIndex = req.PageIndex
	resp.PageNum = req.PageNum
	return resp
}

// GetClassByID : get a certain class
func (m *BackStageManage) GetClassByID(r *http.Request) proto.Message {
	req := &protocol.GetClassByIDReq{}
	resp := &protocol.GetClassByIDResp{Status: &protocol.Status{}}

	status := m.checkArgsandAuth(r, req)
	if status != nil {
		logger.Sugar.Error(status.Message)
		resp.Status = status
		return resp
	}

	class, err := m.db.GetClassByID(req.Id)
	if err != nil {
		logger.Sugar.Errorf("failed to get the class : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = fmt.Sprintf("failed to get the class")
		return resp
	}

	resp.Class = class.TurnProto()
	return resp
}
