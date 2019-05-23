package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/qinhan-shu/gp-server/utils"
)

func TestRandInt(t *testing.T) {
	min := 0
	max := 30
	var randomNum []int
	// regular
	for i := 0; i < 100; i++ {
		random, err := utils.RandInt(min, max)
		if err != nil {
			t.Errorf("failed to get rand : %v", err)
			return
		}
		randomNum = append(randomNum, random)
	}
	t.Logf("%+v", randomNum)

	// error
	_, err := utils.RandInt(max, min)
	if !assert.EqualError(t, err, utils.ErrIllegalRange.Error()) {
		t.Error("Unexpected error result")
		return
	}

}

func Benchmarkrandom(b *testing.B) {
	min := 1
	max := 10000
	for i := 0; i < b.N; i++ {
		utils.RandInt(min, max)
	}
}
