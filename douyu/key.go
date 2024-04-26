package douyu

import (
	"errors"
	"github.com/starskim/DDBOT/lsp/buntdb"
	"strconv"
	"strings"
)

func DouyuGroupConcernStateKey(keys ...interface{}) string {
	return buntdb.NamedKey("DouyuConcernState", keys)
}
func DouyuGroupConcernConfigKey(keys ...interface{}) string {
	return buntdb.NamedKey("DouyuConcernConfig", keys)
}
func DouyuFreshKey(keys ...interface{}) string {
	return buntdb.NamedKey("douyuFresh", keys)
}
func DouyuCurrentLiveKey(keys ...interface{}) string {
	return buntdb.NamedKey("DouyuCurrentLive", keys)
}
func DouyuGroupAtAllMarkKey(keys ...interface{}) string {
	return buntdb.NamedKey("DouyuGroupAtAll", keys)
}

func ParseConcernStateKeyWithInt64(key string) (groupCode int64, id int64, err error) {
	keys := strings.Split(key, ":")
	if len(keys) != 3 {
		return 0, 0, errors.New("invalid key")
	}
	groupCode, err = strconv.ParseInt(keys[1], 10, 64)
	if err != nil {
		return 0, 0, err
	}
	id, err = strconv.ParseInt(keys[2], 10, 64)
	if err != nil {
		return 0, 0, err
	}
	return groupCode, id, nil
}
