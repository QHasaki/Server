package judge

import (
	"net/http"
	"time"

	"github.com/golang/protobuf/proto"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/model/xorm"
	"github.com/qinhan-shu/gp-server/protocol"
	"github.com/qinhan-shu/gp-server/utils"
	"github.com/qinhan-shu/gp-server/utils/judge"
)

// Judge : judge
func (j *Judge) Judge(r *http.Request) proto.Message {
	req := &protocol.JudgeRequest{}
	resp := &protocol.JudgeResponse{Status: &protocol.Status{}}
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
	userID, err := j.cache.GetUserIDByToken(token)
	if err != nil {
		logger.Sugar.Errorf("failed to get token : %v", err)
		resp.Status.Code = protocol.Code_UNAUTHORIZATED
		resp.Status.Message = "invalid token"
		return resp
	}

	problem, err := j.db.GetProblemByID(req.Id)
	if err != nil {
		logger.Sugar.Errorf("failed to get problem by id : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "failed to get problem by id"
		return resp
	}

	client, err := j.getJudgeServer()
	if err != nil {
		logger.Sugar.Errorf("failed to get judge server : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "failed to get judge server"
		return resp
	}

	judgeResp, err := client.JudgeWithRequest(&judge.JudgeRequest{
		Src:            req.Src,
		LanguageConfig: getJudgeConfig(req.Language),
		MaxCpuTime:     int64(problem.JudgeLimitTime),
		MaxMemory:      int64(problem.JudgeLimitMem),
		TestCaseId:     problem.JudgeFile,
	})
	if err != nil {
		logger.Sugar.Errorf("failed to judge : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "failed to judge"
		return resp
	}

	resp.Result = 0
	var submit *model.UserProblem
	for _, data := range judgeResp.SliceData() {
		resp.Results = append(resp.Results, &protocol.JudgeResult{
			JudgeResult: protocol.JudgeResult_Result(data.Result),
			CpuTime:     int64(data.CpuTime),
			RealTime:    int64(data.RealTime),
			Memory:      int64(data.Memory),
			Signal:      int64(data.Signal),
			ExitCode:    int64(data.ExitCode),
		})
		submit = &model.UserProblem{
			Code:            req.Src,
			ProblemId:       req.Id,
			RunningLanguage: int(req.Language),
			RunningMem:      int(data.Memory),
			RunningTime:     int(data.RealTime),
			SubmitTime:      int(time.Now().Unix()),
			Ispass:          1,
			UserId:          userID,
		}
		if data.Result != 0 {
			submit.Ispass = 0
			resp.Result = int64(data.Result)
		}
	}
	if err := j.db.AddSubmitRecord(submit); err != nil {
		logger.Sugar.Errorf("failed to add submit record : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "failed to add submit record"
		return resp
	}

	return resp
}
