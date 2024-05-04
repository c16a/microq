package handlers

import (
	"encoding/json"
	"errors"
	"github.com/c16a/microq/broker"
	"github.com/c16a/microq/events"
	"github.com/c16a/microq/storage"
	"github.com/google/uuid"
	"strings"
)

func handlePublish(message []byte, client *broker.ConnectedClient, broker *broker.Broker, sp storage.Provider) error {

	if !client.IsIdentified() {
		client.WriteInterface(&events.PubAckEvent{
			Kind:    events.PubAck,
			Success: false,
		})
		return errors.New("unidentified client")
	}

	var event events.PubEvent
	err := json.Unmarshal(message, &event)
	if err != nil {
		return err
	}

	topic := event.Topic
	if strings.Contains(topic, "/") || strings.Contains(topic, "\\") || strings.Contains(topic, "..") {
		client.WriteInterface(&events.PubAckEvent{
			Kind:    events.PubAck,
			Success: false,
		})
		return errors.New("topic not allowed")
	}

	event.PacketId = uuid.New().String()

	if event.Retain {
		if err := sp.SaveMessage(&event); err != nil {
			return err
		}
	}
	go broker.Broadcast(event)

	switch event.QoS {
	case 1:
		pubAckEvent := &events.PubAckEvent{
			Kind:     events.PubAck,
			PacketId: event.PacketId,
		}
		client.WriteInterface(pubAckEvent)
		break
	case 2:
		pubRecEvent := &events.PubRecEvent{
			Kind:     events.PubRec,
			PacketId: event.PacketId,
		}
		client.WriteInterface(pubRecEvent)
		break
	}
	return nil
}
