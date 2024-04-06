package handlers

import (
	"encoding/json"
	"github.com/c16a/microq/broker"
	"github.com/c16a/microq/events"
)

func handlePublish(message []byte, broker *broker.Broker) error {
	var event events.PublishEvent
	err := json.Unmarshal(message, &event)
	if err != nil {
		return err
	}
	return broker.Broadcast(event)
}
