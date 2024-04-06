package handlers

import (
	"encoding/json"
	"github.com/c16a/microq/broker"
	"github.com/c16a/microq/events"
)

func handleUnsubscribe(message []byte, client *broker.ConnectedClient) error {
	var event events.UnsubEvent
	err := json.Unmarshal(message, &event)
	if err != nil {
		return err
	}
	client.UnsubscribeFromTopic(event.Topic)
	unsubackEvent := &events.UnsubAckEvent{
		Kind:    events.UnSubAck,
		Success: true,
		Topic:   event.Topic,
	}
	return client.WriteInterface(unsubackEvent)
}
