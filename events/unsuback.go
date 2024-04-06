package events

const UnSubAck = "unsuback"

type UnsubAckEvent struct {
	Kind    string `json:"kind"`
	Success bool   `json:"success"`
	Pattern string `json:"pattern"`
}
