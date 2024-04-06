package handlers

import (
	"encoding/json"
	"github.com/c16a/microq/broker"
	"github.com/c16a/microq/events"
)

func handleSubscribe(message []byte, client *broker.ConnectedClient) error {
	var event events.SubscribeEvent
	err := json.Unmarshal(message, &event)
	if err != nil {
		return err
	}
	client.SubscribeToTopic(event.Topic, event.Group)
	return nil
}
