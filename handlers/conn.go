package handlers

import (
	"encoding/json"
	"github.com/c16a/microq/broker"
	"github.com/c16a/microq/events"
)

func handleConn(message []byte, client *broker.ConnectedClient) error {
	var event events.ConnEvent
	err := json.Unmarshal(message, &event)
	if err != nil {
		return err
	}

	client.SetId(event.ClientId)
	subackEvent := &events.ConnAckEvent{
		Kind:     events.ConnAck,
		Status:   true,
		ClientId: event.ClientId,
	}
	return client.WriteInterface(subackEvent)
}
