package handlers

import (
	"encoding/json"
	"errors"
	"github.com/c16a/microq/broker"
	"github.com/c16a/microq/events"
)

func handleSubscribe(message []byte, client *broker.ConnectedClient) error {
	if !client.IsIdentified() {
		client.WriteInterface(&events.SubAckEvent{
			Kind:    events.SubAck,
			Success: false,
		})
		return errors.New("unidentified client")
	}

	var event events.SubEvent
	err := json.Unmarshal(message, &event)
	if err != nil {
		return err
	}
	client.SubscribeToPattern(event.Pattern, event.Group)
	subackEvent := &events.SubAckEvent{
		Kind:    events.SubAck,
		Success: true,
		Pattern: event.Pattern,
	}
	return client.WriteInterface(subackEvent)
}
