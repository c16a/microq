package events

const PubRec = "pubrec"

type PubRecEvent struct {
	Kind     string `json:"kind"`
	PacketId string `json:"packet_id"`
}
