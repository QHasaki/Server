package db_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/qinhan-shu/gp-server/model/transform"
	"github.com/qinhan-shu/gp-server/model/xorm"
)

func TestMysqlDriver_GetProblemsNum(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	num, err := mysqlDriver.GetProblemsNum()
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("the num of problems : %d", num)
}
func TestMysqlDriver_GetProblems(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	var pageIndex int64 = 1
	var pageNum int64 = 3
	problems, err := mysqlDriver.GetProblems(pageNum, pageIndex)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(len(problems))
	for _, problem := range problems {
		t.Logf("%+v\n", problem.Detail)
		t.Logf("%+v\n", problem.InAndOutExample)
	}
}

func TestMysqlDriver_GetProblemsByTagID(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	var tagID = 2
	var pageIndex int64 = 1
	var pageNum int64 = 1000
	problems, err := mysqlDriver.GetProblemsByTagID(pageNum, pageIndex, tagID)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(len(problems))
	for _, problem := range problems {
		t.Logf("%+v\n", problem.Detail.Id)
		t.Logf("%+v\n", problem.Detail.Tags)
		t.Logf("%+v\n", problem.InAndOutExample)
	}
}

func TestMysqlDriver_GetProblemByID(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	var id int64 = 2
	problem, err := mysqlDriver.GetProblemByID(id)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v\n", problem.Detail)
	t.Logf("%+v\n", problem.InAndOutExample)
}

func TestMysqlDriver_AddProblem(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	data, _ := json.Marshal([]int64{1, 2, 3})
	problem := &model.Problem{
		Title:          "求和问题",
		Description:    "求两个数的和",
		InDescription:  "输入两个int型数",
		OutDescription: "输出两个数的和",
		Hint:           "无提示",
		JudgeLimitTime: 10,
		JudgeLimitMem:  10,
		Difficulty:     1,
		Tags:           string(data),
	}
	testData := []*model.TestData{
		&model.TestData{
			In:  "in 1",
			Out: "out 1",
		},
		&model.TestData{
			In:  "in 2",
			Out: "out 2",
		},
	}
	if err := mysqlDriver.AddProblem(&transform.IntactProblem{
		Detail:          problem,
		InAndOutExample: testData,
	}); err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v\n", problem)
}

func TestMysqlDriver_UpdateProblem(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	var problemID int64 = 12
	originProblem, err := mysqlDriver.GetProblemByID(problemID)
	if err != nil {
		t.Error(err)
		return
	}

	change := &transform.IntactProblem{
		Detail: &model.Problem{
			Id:    problemID,
			Title: originProblem.Detail.Title + "000",
		},
	}

	if err := mysqlDriver.UpdateProblem(change); err != nil {
		t.Error(err)
		return
	}

	changedProblem, err := mysqlDriver.GetProblemByID(problemID)
	if err != nil {
		t.Error(err)
		return
	}

	if !assert.NotEqual(t, originProblem.Detail.Title, changedProblem.Detail.Title) {
		t.Error("filed [Title] is not changed")
		return
	}

	if !assert.Equal(t, changedProblem.Detail.Title, change.Detail.Title) {
		t.Error("filed [Title] is not changed to expected value")
		return
	}

	if !assert.Equal(t, originProblem.Detail.Description, changedProblem.Detail.Description) {
		t.Error("other filed [Description] is changed")
		return
	}
}

func TestMysqlDriver_GetAllProblems(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	problems, err := mysqlDriver.GetAllProblems()
	if err != nil {
		t.Error(err)
		return
	}
	for _, problem := range problems {
		t.Log(problem)
	}
}

func TestAddSomeProblems(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}
	num := 10
	data, _ := json.Marshal([]int64{1, 2, 3})
	for i := 0; i < num; i++ {
		problem := &model.Problem{
			Title:          "求和问题",
			Description:    "求两个数的和",
			InDescription:  "输入两个int型数",
			OutDescription: "输出两个数的和",
			Hint:           "无提示",
			JudgeLimitTime: 10,
			JudgeLimitMem:  10,
			Difficulty:     1,
			Tags:           string(data),
		}
		testData := []*model.TestData{
			&model.TestData{
				In:  "in 1",
				Out: "out 1",
			},
			&model.TestData{
				In:  "in 2",
				Out: "out 2",
			},
		}
		if err := mysqlDriver.AddProblem(&transform.IntactProblem{
			Detail:          problem,
			InAndOutExample: testData,
		}); err != nil {
			t.Error(err)
			return
		}

		t.Logf("%+v\n", problem)
	}
}
