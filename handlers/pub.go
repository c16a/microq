package handlers

import (
	"encoding/json"
	"github.com/c16a/microq/broker"
	"github.com/c16a/microq/events"
	"github.com/google/uuid"
)

func handlePublish(message []byte, client *broker.ConnectedClient, broker *broker.Broker) error {
	var event events.PubEvent
	err := json.Unmarshal(message, &event)
	if err != nil {
		return err
	}
	if event.QoS == 1 {
		pubAckEvent := &events.PubAckEvent{
			Kind:     events.PubAck,
			PacketId: uuid.NewString(),
		}
		client.WriteInterface(pubAckEvent)
	}
	if event.QoS == 2 {
		pubRecEvent := &events.PubRecEvent{
			Kind:     events.PubRec,
			PacketId: uuid.NewString(),
		}
		client.WriteInterface(pubRecEvent)
	}
	go broker.Broadcast(event)
	return nil
}
