package events

const Unsub = "unsub"

type UnsubEvent struct {
	Kind    string `json:"kind"`
	Pattern string `json:"pattern"`
}
