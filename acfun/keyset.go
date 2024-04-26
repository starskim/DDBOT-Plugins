package acfun

type extraKey struct{}

func (e *extraKey) UserInfoKey(keys ...interface{}) string {
	return AcfunUserInfoKey(keys...)
}

func (e *extraKey) LiveInfoKey(keys ...interface{}) string {
	return AcfunLiveInfoKey(keys...)
}

func (e *extraKey) NotLiveKey(keys ...interface{}) string {
	return AcfunNotLiveKey(keys...)
}

func (e *extraKey) UidFirstTimestamp(keys ...interface{}) string {
	return AcfunUidFirstTimestampKey(keys...)
}
