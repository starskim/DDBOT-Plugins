package youtube

import (
	"errors"
	"github.com/starskim/DDBOT/lsp/buntdb"
	"strconv"
	"strings"
)

func YoutubeGroupConcernStateKey(keys ...interface{}) string {
	return buntdb.NamedKey("YoutubeConcernState", keys)
}
func YoutubeGroupConcernConfigKey(keys ...interface{}) string {
	return buntdb.NamedKey("YoutubeConcernConfig", keys)
}
func YoutubeFreshKey(keys ...interface{}) string {
	return buntdb.NamedKey("youtubeFresh", keys)
}
func YoutubeUserInfoKey(keys ...interface{}) string {
	return buntdb.NamedKey("YoutubeUserInfo", keys)
}
func YoutubeInfoKey(keys ...interface{}) string {
	return buntdb.NamedKey("YoutubeInfo", keys)
}
func YoutubeVideoKey(keys ...interface{}) string {
	return buntdb.NamedKey("YoutubeVideo", keys)
}
func YoutubeGroupAtAllMarkKey(keys ...interface{}) string {
	return buntdb.NamedKey("YoutubeGroupAtAll", keys)
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
