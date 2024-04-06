package events

const Unsub = "unsub"

type UnsubEvent struct {
	Kind  string `json:"kind"`
	Topic string `json:"topic"`
}
