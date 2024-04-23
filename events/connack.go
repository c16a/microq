package events

const ConnAck = "connack"

type ConnAckEvent struct {
	Kind     string `json:"kind"`
	ClientId string `json:"client_id"`
	Status   bool   `json:"status"`
}
