package events

const Conn = "conn"

type ConnEvent struct {
	Kind     string `json:"kind"`
	ClientId string `json:"client_id"`
}
