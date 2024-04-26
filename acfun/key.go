package acfun

import "github.com/starskim/DDBOT/lsp/buntdb"

func AcfunUserInfoKey(keys ...interface{}) string {
	return buntdb.NamedKey("AcfunUserInfo", keys)
}
func AcfunLiveInfoKey(keys ...interface{}) string {
	return buntdb.NamedKey("AcfunLiveInfo", keys)
}
func AcfunNotLiveKey(keys ...interface{}) string {
	return buntdb.NamedKey("AcfunNotLiveCount", keys)
}
func AcfunUidFirstTimestampKey(keys ...interface{}) string {
	return buntdb.NamedKey("AcfunUidFirstTimestamp", keys)
}
