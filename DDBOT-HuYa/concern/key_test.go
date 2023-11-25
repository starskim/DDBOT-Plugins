package huya

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	GroupCode1 int64 = 123456
	Uid        int64 = 777
	Sid              = "uid"
)

func TestParseConcernStateKeyWithString(t *testing.T) {
	var testCase = []string{
		HuyaGroupConcernStateKey(GroupCode1, Sid),
	}
	var expected = [][]interface{}{
		{
			GroupCode1, Sid,
		},
	}
	assert.Equal(t, len(expected), len(testCase))
	for index := range testCase {
		a, b, err := ParseConcernStateKeyWithString(testCase[index])
		assert.Nil(t, err)
		assert.EqualValues(t, []interface{}{a, b}, expected[index])
	}
}
func TestKeys(t *testing.T) {
	HuyaGroupConcernStateKey()
	HuyaGroupConcernConfigKey()
	HuyaFreshKey()
	HuyaCurrentLiveKey()
	HuyaGroupAtAllMarkKey()
}
