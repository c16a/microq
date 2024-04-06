package handlers

import (
	"encoding/json"
	"github.com/c16a/microq/broker"
	"github.com/c16a/microq/events"
	"testing"
)

func TestHandleSubscribe(t *testing.T) {
	client := broker.NewConnectedClient(&TestingWebSocketConnection{}, "client-1")

	event := &events.SubEvent{
		Kind:  events.Sub,
		Group: "g1",
		Topic: "t1",
	}
	message, err := json.Marshal(event)
	if err != nil {
		t.Fatal(err)
	}

	err = handleSubscribe(message, client)

	subscription := client.GetSubscription("t1")
	if subscription == nil {
		t.Fatal("subscription was nil")
	}
	if !subscription.IsActive() {
		t.Fatal("subscription was not active")
	}
	if subscription.GetGroup() != "g1" {
		t.Fatal("subscription group was wrong")
	}
}
