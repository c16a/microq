package events

const Sub = "sub"

type SubEvent struct {
	Kind    string `json:"kind"`
	Group   string `json:"group"`
	Pattern string `json:"pattern"`
}
