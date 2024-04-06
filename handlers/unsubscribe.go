package handlers

import (
	"encoding/json"
	"github.com/c16a/microq/broker"
	"github.com/c16a/microq/events"
)

func handleUnsubscribe(message []byte, client *broker.ConnectedClient) error {
	var event events.UnsubscribeEvent
	err := json.Unmarshal(message, &event)
	if err != nil {
		return err
	}
	client.UnsubscribeFromTopics(event.Topics)
	return nil
}
