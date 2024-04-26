package weibo

type extraKeySet struct{}

func (*extraKeySet) UserInfoKey(keys ...interface{}) string {
	return WeiboUserInfoKey(keys...)
}

func (*extraKeySet) NewsInfoKey(keys ...interface{}) string {
	return WeiboNewsInfoKey(keys...)
}

func (*extraKeySet) MarkMblogIdKey(keys ...interface{}) string {
	return WeiboMarkMblogIdKey(keys...)
}
