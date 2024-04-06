package broker

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"sync"
)

type ConnectedClient struct {
	id            string
	subscriptions []*Subscription
	conn          WebSocketConnection
	mutex         sync.RWMutex
}

func NewConnectedClient(conn WebSocketConnection, id string) *ConnectedClient {
	return &ConnectedClient{
		id:            id,
		conn:          conn,
		mutex:         sync.RWMutex{},
		subscriptions: make([]*Subscription, 0),
	}
}

func (client *ConnectedClient) WriteDataMessage(data []byte) error {
	return client.conn.WriteMessage(websocket.TextMessage, data)
}

func (client *ConnectedClient) WriteInterface(v any) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return client.WriteDataMessage(data)
}

func (client *ConnectedClient) GetSubscription(topic string) *Subscription {
	client.mutex.RLock()
	defer client.mutex.RUnlock()

	for _, sub := range client.subscriptions {
		if sub.Matches(topic) {
			return sub
		}
	}

	return nil
}

func (client *ConnectedClient) SubscribeToPattern(pattern string, group string) {
	client.mutex.Lock()
	defer client.mutex.Unlock()

	subscription := &Subscription{active: true, pattern: pattern}
	if len(group) > 0 {
		subscription.group = group
	}

	client.subscriptions = append(client.subscriptions, subscription)
}

func (client *ConnectedClient) UnsubscribeFromTopic(topic string) {
	client.mutex.Lock()
	defer client.mutex.Unlock()

	for i, subscription := range client.subscriptions {
		if subscription.Matches(topic) {
			client.subscriptions = append(client.subscriptions[:i], client.subscriptions[i+1:]...)
		}
	}
}
