package events

const Unsubscribe = "unsubscribe"

type UnsubscribeEvent struct {
	Kind   string   `json:"kind"`
	Topics []string `json:"topics"`
}
