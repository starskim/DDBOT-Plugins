package test

import (
	localdb "github.com/starskim/DDBOT/lsp/buntdb"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 这个包只允许在单元测试中使用

const (
	UID1 int64 = 777
	G1   int64 = 123456
	G2   int64 = 654321

	NAME1 = "name1"
	NAME2 = "name2"
)

func InitBuntdb(t *testing.T) {
	assert.Nil(t, localdb.InitBuntDB(localdb.MEMORYDB))
}
func CloseBuntdb(t *testing.T) {
	assert.Nil(t, localdb.Close())
}
