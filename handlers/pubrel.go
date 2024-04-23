package handlers

import (
	"encoding/json"
	"errors"
	"github.com/c16a/microq/broker"
	"github.com/c16a/microq/events"
)

func handlePubrel(message []byte, client *broker.ConnectedClient) error {
	if !client.IsIdentified() {
		client.WriteInterface(&events.PubCompEvent{
			Kind:    events.PubComp,
			Success: false,
		})
		return errors.New("unidentified client")
	}

	var event events.PubRelEvent
	err := json.Unmarshal(message, &event)
	if err != nil {
		return err
	}
	pubCompEvent := &events.PubCompEvent{
		Kind:     events.PubComp,
		PacketId: event.PacketId,
	}
	client.WriteInterface(pubCompEvent)
	return nil
}
