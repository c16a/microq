package events

const ConnAck = "connack"

type ConnAckEvent struct {
	Kind     string `json:"kind"`
	ClientId string `json:"client_id"`
	Success  bool   `json:"success"`
}
