package handlers

import (
	"github.com/c16a/microq/broker"
	"github.com/c16a/microq/events"
	"github.com/c16a/microq/storage"
)

func HandleMessage(client *broker.ConnectedClient, broker *broker.Broker, sp storage.Provider, message []byte) error {
	kind := events.GetKindFromJson(message)
	switch kind {
	case events.Pub:
		return handlePublish(message, client, broker, sp)
	case events.PubRel:
		return handlePubrel(message, client)
	case events.Sub:
		return handleSubscribe(message, client)
	case events.Unsub:
		return handleUnsubscribe(message, client)
	case events.Conn:
		return handleConn(message, client, broker)
	}
	return nil
}
