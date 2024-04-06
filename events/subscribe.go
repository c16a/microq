package events

const Subscribe = "subscribe"

type SubscribeEvent struct {
	Kind  string `json:"kind"`
	Group string `json:"group"`
	Topic string `json:"topic"`
}
