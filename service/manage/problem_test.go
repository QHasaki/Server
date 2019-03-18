package manage_test

import (
	"testing"

	"github.com/golang/protobuf/proto"

	"github.com/qinhan-shu/gp-server/module"
	"github.com/qinhan-shu/gp-server/protocol"
	"github.com/qinhan-shu/gp-server/service/config"
	"github.com/qinhan-shu/gp-server/service/manage"
	"github.com/qinhan-shu/gp-server/utils/mode"
)

func TestUserManage_AddProblem(t *testing.T) {
	mode.SetMode(mode.TestMode)
	dataStorage, err := config.NewConfig().GetDataStorage()
	if err != nil {
		t.Error(err)
		return
	}
	managerModule := manage.NewBackStageManager(dataStorage)

	// get user operations
	reqBytes, err := proto.Marshal(&protocol.AddProblemReq{
		Problem: &protocol.Problem{
			Title:       "这是一个测试题目(求和)",
			Description: "题目描述",
			In:          "输入描述",
			Out:         "输出描述",
			Hint:        "没有提示",
			InOutExamples: []*protocol.ProblemExample{
				&protocol.ProblemExample{
					Input:  "1 1",
					Output: "2",
				},
				&protocol.ProblemExample{
					Input:  "2 2",
					Output: "4",
				},
			},
			JudgeLimit: &protocol.ProblemJudgeLimit{
				Time: "62s",
				Mem:  "7m",
			},
			Tags:       []string{"求和", "数组", "树"},
			Difficluty: protocol.ProblemDifficluty_HARD,
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

	data := managerModule.AddProblem(args)
	resp := data.(*protocol.AddProblemResp)
	if resp.Code != protocol.Code_OK {
		t.Error(err)
		return
	}
}
