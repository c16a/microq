package events

import "github.com/valyala/fastjson"

// This doesn't parse the entire JSON, so it's (probably) fast
func GetKindFromJson(data []byte) string {
	return fastjson.GetString(data, "kind")
}
