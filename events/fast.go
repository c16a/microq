package events

import (
	"encoding/json"
)

type kindOnly struct {
	Kind string `json:"kind"`
}

// This doesn't parse the entire JSON, so it's (probably) fast
func GetKindFromJson(data []byte) string {
	var k kindOnly
	err := json.Unmarshal(data, &k)
	if err != nil {
		return ""
	}
	return k.Kind
}
