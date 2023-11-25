package bilibili

import (
	"errors"
	"github.com/starskim/DDBOT/lsp/buntdb"
	"strconv"
	"strings"
)

func BilibiliGroupConcernStateKey(keys ...interface{}) string {
	return buntdb.NamedKey("ConcernState", keys)
}
func BilibiliGroupConcernConfigKey(keys ...interface{}) string {
	return buntdb.NamedKey("ConcernConfig", keys)
}
func BilibliFreshKey(keys ...interface{}) string {
	return buntdb.NamedKey("fresh", keys)
}
func BilibiliCurrentLiveKey(keys ...interface{}) string {
	return buntdb.NamedKey("CurrentLive", keys)
}
func BilibiliCurrentNewsKey(keys ...interface{}) string {
	return buntdb.NamedKey("CurrentNews", keys)
}
func BilibiliDynamicIdKey(keys ...interface{}) string {
	return buntdb.NamedKey("DynamicId", keys)
}
func BilibiliUidFirstTimestampKey(keys ...interface{}) string {
	return buntdb.NamedKey("UidFirstTimestamp", keys)
}
func BilibiliUserCookieInfoKey(keys ...interface{}) string {
	return buntdb.NamedKey("UserCookieInfo", keys)
}
func BilibiliNotLiveCountKey(keys ...interface{}) string {
	return buntdb.NamedKey("NotLiveCount", keys)
}
func BilibiliUserInfoKey(keys ...interface{}) string {
	return buntdb.NamedKey("UserInfo", keys)
}
func BilibiliUserStatKey(keys ...interface{}) string {
	return buntdb.NamedKey("UserStat", keys)
}
func BilibiliGroupAtAllMarkKey(keys ...interface{}) string {
	return buntdb.NamedKey("GroupAtAll", keys)
}
func BilibiliCompactMarkKey(keys ...interface{}) string {
	return buntdb.NamedKey("CompactMark", keys)
}
func BilibiliNotifyMsgKey(keys ...interface{}) string {
	return buntdb.NamedKey("NotifyMsg", keys)
}
func BilibiliActiveTimestampKey(keys ...interface{}) string {
	return buntdb.NamedKey("ActiveTimestamp", keys)
}
func BilibiliLastFreshKey(keys ...interface{}) string {
	return buntdb.NamedKey("BilibiliLastFresh", keys)
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
