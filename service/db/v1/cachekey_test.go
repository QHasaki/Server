package db_test

import (
	"testing"

	"github.com/qinhan-shu/gp-server/module/v1"
	"github.com/qinhan-shu/gp-server/service/db/v1"
)

func TestMakeCacheKey(t *testing.T) {
	document := "player"
	where := make(module.Data)
	where["id"] = 1
	where["nickname"] = "qinhan"

	t.Log(db.MakeCacheKey(document, where))
}
