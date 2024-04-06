package events

const PubRel = "pubrel"

type PubRelEvent struct {
	Kind     string `json:"kind"`
	PacketId string `json:"packet_id"`
}
