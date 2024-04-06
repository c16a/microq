package handlers

import (
	"encoding/json"
	"github.com/c16a/microq/broker"
	"github.com/c16a/microq/events"
)

func handlePubrel(message []byte, client *broker.ConnectedClient) error {
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
