package test

import (
	"fmt"
	localdb "github.com/starskim/DDBOT/lsp/buntdb"
	"github.com/starskim/DDBOT/lsp/concern_type"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 这个包只允许在单元测试中使用

const (
	ROOMID1    int64 = 1
	UID1       int64 = 777
	UID2       int64 = 778
	DynamicID1 int64 = 1001
	G1         int64 = 123456
	G2         int64 = 654321
	TIMESTAMP1 int64 = 1624126814
	TIMESTAMP2 int64 = 1624126914

	NAME1 = "name1"
	NAME2 = "name2"

	BVID1 = "bvid1"
	BVID2 = "bvid2"

	ID1 = 2001
)

const (
	BibiliLive concern_type.Type = "bilibiliLive"
)

func InitBuntdb(t *testing.T) {
	assert.Nil(t, localdb.InitBuntDB(localdb.MEMORYDB))
}
func CloseBuntdb(t *testing.T) {
	assert.Nil(t, localdb.Close())
}

func FakeImage(size int) string {
	if size == 0 {
		size = 150
	}
	return fmt.Sprintf("https://via.placeholder.com/%v", size)
}
