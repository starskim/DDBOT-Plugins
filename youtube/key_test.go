package youtube

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	GroupCode1 int64 = 123456
	Sid              = "uid"
)

func TestParseConcernStateKeyWithInt642(t *testing.T) {
	var testCase = []string{
		"wrong_key",
		YoutubeGroupConcernStateKey(GroupCode1, Sid),
	}

	for _, key := range testCase {
		_, _, err := ParseConcernStateKeyWithInt64(key)
		assert.NotNil(t, err)
	}

}

func TestParseConcernStateKeyWithString(t *testing.T) {
	var testCase = []string{
		YoutubeGroupConcernStateKey(GroupCode1, Sid),
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

func TestParseConcernStateKeyWithString2(t *testing.T) {
	var testCase = []string{
		"wrong_key",
		YoutubeGroupConcernStateKey("wrong_group", Sid),
	}
	for _, key := range testCase {
		_, _, err := ParseConcernStateKeyWithString(key)
		assert.NotNil(t, err)
	}
}

func TestKeys(t *testing.T) {
	YoutubeGroupConcernConfigKey()
	YoutubeFreshKey()
	YoutubeUserInfoKey()
	YoutubeInfoKey()
	YoutubeVideoKey()
	YoutubeGroupAtAllMarkKey()

}
