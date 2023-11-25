package douyu

type keySet struct {
}

func (l *keySet) GroupAtAllMarkKey(keys ...interface{}) string {
	return DouyuGroupAtAllMarkKey(keys...)
}

func (l *keySet) GroupConcernConfigKey(keys ...interface{}) string {
	return DouyuGroupConcernConfigKey(keys...)
}

func (l *keySet) GroupConcernStateKey(keys ...interface{}) string {
	return DouyuGroupConcernStateKey(keys...)
}

func (l *keySet) FreshKey(keys ...interface{}) string {
	return DouyuFreshKey(keys...)
}

func (l *keySet) ParseGroupConcernStateKey(key string) (int64, interface{}, error) {
	return ParseConcernStateKeyWithInt64(key)
}

type extraKey struct {
}

func (l *extraKey) CurrentLiveKey(keys ...interface{}) string {
	return DouyuCurrentLiveKey(keys...)
}

func NewExtraKey() *extraKey {
	return &extraKey{}
}

func NewKeySet() *keySet {
	return &keySet{}
}
