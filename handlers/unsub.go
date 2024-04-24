package handlers

import (
	"encoding/json"
	"errors"
	"github.com/c16a/microq/broker"
	"github.com/c16a/microq/events"
)

func handleUnsubscribe(message []byte, client *broker.ConnectedClient) error {
	if !client.IsIdentified() {
		client.WriteInterface(&events.UnsubAckEvent{
			Kind:    events.PubAck,
			Success: false,
		})
		return errors.New("unidentified client")
	}

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
