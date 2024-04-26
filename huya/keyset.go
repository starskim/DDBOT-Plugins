package huya

type keySet struct {
}

func (l *keySet) GroupAtAllMarkKey(keys ...interface{}) string {
	return HuyaGroupAtAllMarkKey(keys...)
}

func (l *keySet) GroupConcernConfigKey(keys ...interface{}) string {
	return HuyaGroupConcernConfigKey(keys...)
}

func (l *keySet) GroupConcernStateKey(keys ...interface{}) string {
	return HuyaGroupConcernStateKey(keys...)
}

func (l *keySet) FreshKey(keys ...interface{}) string {
	return HuyaFreshKey(keys...)
}

func (l *keySet) ParseGroupConcernStateKey(key string) (int64, interface{}, error) {
	return ParseConcernStateKeyWithString(key)
}

type extraKey struct{}

func (k extraKey) CurrentLiveKey(keys ...interface{}) string {
	return HuyaCurrentLiveKey(keys...)
}

func NewExtraKey() *extraKey {
	return &extraKey{}
}

func NewKeySet() *keySet {
	return &keySet{}
}
