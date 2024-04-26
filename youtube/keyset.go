package youtube

type KeySet struct {
}

func (k *KeySet) GroupAtAllMarkKey(keys ...interface{}) string {
	return YoutubeGroupAtAllMarkKey(keys...)
}

func (k *KeySet) GroupConcernConfigKey(keys ...interface{}) string {
	return YoutubeGroupConcernConfigKey(keys...)
}

func (k *KeySet) GroupConcernStateKey(keys ...interface{}) string {
	return YoutubeGroupConcernStateKey(keys...)
}

func (k *KeySet) FreshKey(keys ...interface{}) string {
	return YoutubeFreshKey(keys...)
}

func (k *KeySet) ParseGroupConcernStateKey(key string) (int64, interface{}, error) {
	return ParseConcernStateKeyWithString(key)
}

func NewKeySet() *KeySet {
	return new(KeySet)
}

type extraKey struct {
}

func (e *extraKey) UserInfoKey(keys ...interface{}) string {
	return YoutubeUserInfoKey(keys...)
}

func (e *extraKey) InfoKey(keys ...interface{}) string {
	return YoutubeInfoKey(keys...)
}

func (e *extraKey) VideoKey(keys ...interface{}) string {
	return YoutubeVideoKey(keys...)
}

func NewExtraKey() *extraKey {
	return &extraKey{}
}
