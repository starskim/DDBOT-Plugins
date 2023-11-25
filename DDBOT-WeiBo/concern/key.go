package weibo

import "github.com/starskim/DDBOT/lsp/buntdb"

func WeiboUserInfoKey(keys ...interface{}) string {
	return buntdb.NamedKey("WeiboUserInfo", keys)
}
func WeiboNewsInfoKey(keys ...interface{}) string {
	return buntdb.NamedKey("WeiboNewsInfo", keys)
}
func WeiboMarkMblogIdKey(keys ...interface{}) string {
	return buntdb.NamedKey("WeiboMarkMblogId", keys)
}
