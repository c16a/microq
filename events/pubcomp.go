package events

const PubComp = "pubcomp"

type PubCompEvent struct {
	Kind     string `json:"kind"`
	PacketId string `json:"packet_id"`
}
