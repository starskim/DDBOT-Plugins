package bilibili

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
		BilibiliGroupConcernStateKey(GroupCode1, Uid),
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

func TestParseConcernStateKeyWithInt642(t *testing.T) {
	var testCase = []string{
		"wrong_key",
		BilibiliGroupConcernStateKey("wrong_group", Uid),
	}

	for _, key := range testCase {
		_, _, err := ParseConcernStateKeyWithInt64(key)
		assert.NotNil(t, err)
	}

}
func TestKeys(t *testing.T) {
	BilibiliGroupConcernStateKey()
	BilibiliGroupConcernConfigKey()
	BilibliFreshKey()
	BilibiliCurrentLiveKey()
	BilibiliCurrentNewsKey()
	BilibiliDynamicIdKey()
	BilibiliUidFirstTimestampKey()
	BilibiliUserCookieInfoKey()
	BilibiliNotLiveCountKey()
	BilibiliUserInfoKey()
	BilibiliUserStatKey()
	BilibiliGroupAtAllMarkKey()
	BilibiliNotifyMsgKey()
	BilibiliCompactMarkKey()
	BilibiliLastFreshKey()
	assert.Panics(t, func() {
		BilibiliGroupConcernStateKey(&struct{}{})
	})
}
