package handlers

import (
	"encoding/json"
	"github.com/c16a/microq/broker"
	"github.com/c16a/microq/events"
)

func handleSubscribe(message []byte, client *broker.ConnectedClient) error {
	var event events.SubEvent
	err := json.Unmarshal(message, &event)
	if err != nil {
		return err
	}
	client.SubscribeToPattern(event.Pattern, event.Group)
	subackEvent := &events.SubAckEvent{
		Kind:    events.SubAck,
		Success: true,
		Topic:   event.Pattern,
	}
	return client.WriteInterface(subackEvent)
}
