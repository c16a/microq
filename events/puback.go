package events

const PubAck = "puback"

type PubAckEvent struct {
	Kind     string `json:"kind"`
	PacketId string `json:"packet_id"`
	Success  bool   `json:"success"`
}
