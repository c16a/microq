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
	client.UnsubscribeFromPattern(event.Pattern)
	unsubackEvent := &events.UnsubAckEvent{
		Kind:    events.UnSubAck,
		Success: true,
		Pattern: event.Pattern,
	}
	return client.WriteInterface(unsubackEvent)
}
