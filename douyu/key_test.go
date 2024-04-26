package douyu

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	GroupCode1 int64 = 123456
	Uid        int64 = 777
)

func TestParseConcernStateKeyWithInt64(t *testing.T) {
	var testCase = []string{
		DouyuGroupConcernStateKey(GroupCode1, Uid),
	}
	var expected = [][]int64{
		{
			GroupCode1, Uid,
		},
	}
	assert.Equal(t, len(expected), len(testCase))
	for index := range testCase {
		a, b, err := ParseConcernStateKeyWithInt64(testCase[index])
		assert.Nil(t, err)
		assert.EqualValues(t, []int64{a, b}, expected[index])
	}
}

func TestKeys(t *testing.T) {
	DouyuGroupConcernStateKey()
	DouyuGroupConcernConfigKey()
	DouyuFreshKey()
	DouyuCurrentLiveKey()
	DouyuGroupAtAllMarkKey()
}
