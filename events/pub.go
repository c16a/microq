package events

const Pub = "pub"

type PubEvent struct {
	Kind    string `json:"kind"`
	Message string `json:"message"`
	Topic   string `json:"topic"`
	QoS     int    `json:"qos"`
}
