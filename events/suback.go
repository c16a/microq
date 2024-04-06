package events

const SubAck = "suback"

type SubAckEvent struct {
	Kind    string `json:"kind"`
	Success bool   `json:"success"`
	Topic   string `json:"topic"`
}
