package handlers

import (
	"github.com/c16a/microq/broker"
	"github.com/c16a/microq/events"
)

func HandleMessage(client *broker.ConnectedClient, broker *broker.Broker, message []byte) error {
	kind := events.GetKindFromJson(message)
	switch kind {
	case events.Publish:
		return handlePublish(message, broker)
	case events.Subscribe:
		return handleSubscribe(message, client)
	case events.Unsubscribe:
		return handleUnsubscribe(message, client)
	}
	return nil
}
