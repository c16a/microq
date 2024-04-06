package handlers

import (
	"github.com/c16a/microq/broker"
	"github.com/c16a/microq/events"
)

func HandleMessage(client *broker.ConnectedClient, broker *broker.Broker, message []byte) error {
	kind := events.GetKindFromJson(message)
	switch kind {
	case events.Pub:
		return handlePublish(message, client, broker)
	case events.Sub:
		return handleSubscribe(message, client)
	case events.Unsub:
		return handleUnsubscribe(message, client)
	}
	return nil
}
