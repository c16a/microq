package events

const Publish = "publish"

type PublishEvent struct {
	Kind    string `json:"kind"`
	Message string `json:"message"`
	Topic   string `json:"topic"`
}
