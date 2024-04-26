package huya

import (
	"errors"
	"github.com/starskim/DDBOT/lsp/buntdb"
	"strconv"
	"strings"
)

func HuyaGroupConcernStateKey(keys ...interface{}) string {
	return buntdb.NamedKey("HuyaConcernState", keys)
}
func HuyaGroupConcernConfigKey(keys ...interface{}) string {
	return buntdb.NamedKey("HuyaConcernConfig", keys)
}
func HuyaFreshKey(keys ...interface{}) string {
	return buntdb.NamedKey("huyaFresh", keys)
}
func HuyaCurrentLiveKey(keys ...interface{}) string {
	return buntdb.NamedKey("HuyaCurrentLive", keys)
}
func HuyaGroupAtAllMarkKey(keys ...interface{}) string {
	return buntdb.NamedKey("HuyaGroupAtAll", keys)
}

func ParseConcernStateKeyWithString(key string) (groupCode int64, id string, err error) {
	keys := strings.Split(key, ":")
	if len(keys) != 3 {
		return 0, "", errors.New("invalid key")
	}
	groupCode, err = strconv.ParseInt(keys[1], 10, 64)
	if err != nil {
		return 0, "", err
	}
	return groupCode, keys[2], nil

}
