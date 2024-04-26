package bilibili

type keySet struct {
}

func (k *keySet) GroupAtAllMarkKey(keys ...interface{}) string {
	return BilibiliGroupAtAllMarkKey(keys...)
}

func (k *keySet) GroupConcernConfigKey(keys ...interface{}) string {
	return BilibiliGroupConcernConfigKey(keys...)
}

func (k *keySet) GroupConcernStateKey(keys ...interface{}) string {
	return BilibiliGroupConcernStateKey(keys...)
}

func (k *keySet) FreshKey(keys ...interface{}) string {
	return BilibliFreshKey(keys...)
}

func (k *keySet) ParseGroupConcernStateKey(key string) (int64, interface{}, error) {
	return ParseConcernStateKeyWithInt64(key)
}

type extraKey struct {
}

func (k *extraKey) CurrentLiveKey(keys ...interface{}) string {
	return BilibiliCurrentLiveKey(keys...)
}

func (k *extraKey) UserInfoKey(keys ...interface{}) string {
	return BilibiliUserInfoKey(keys...)
}
func (k *extraKey) UserStatKey(keys ...interface{}) string {
	return BilibiliUserStatKey(keys...)
}
func (k *extraKey) CurrentNewsKey(keys ...interface{}) string {
	return BilibiliCurrentNewsKey(keys...)
}

func (k *extraKey) DynamicIdKey(keys ...interface{}) string {
	return BilibiliDynamicIdKey(keys...)
}

func (k *extraKey) UidFirstTimestamp(keys ...interface{}) string {
	return BilibiliUidFirstTimestampKey(keys...)
}

func (k *extraKey) NotLiveKey(keys ...interface{}) string {
	return BilibiliNotLiveCountKey(keys...)
}

func (k *extraKey) LastFreshKey(keys ...interface{}) string {
	return BilibiliLastFreshKey(keys...)
}

func (k *extraKey) CompactMarkKey(keys ...interface{}) string {
	return BilibiliCompactMarkKey(keys...)
}

func (k *extraKey) NotifyMsgKey(keys ...interface{}) string {
	return BilibiliNotifyMsgKey(keys...)
}

func (k *extraKey) ActiveTimestampKey(keys ...interface{}) string {
	return BilibiliActiveTimestampKey(keys...)
}

func NewKeySet() *keySet {
	return &keySet{}
}
func NewExtraKey() *extraKey {
	return &extraKey{}
}
