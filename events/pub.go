package events

const Pub = "pub"

type PubEvent struct {
	Kind     string `json:"kind"`
	PacketId string `json:"packet_id"`
	Message  string `json:"message"`
	Topic    string `json:"topic"`
	QoS      int    `json:"qos"`
	Retain   bool   `json:"retain"`
}
