package broker

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"sync"
)

type ConnectedClient struct {
	id            string
	subscriptions map[string]*Subscription
	conn          WebSocketConnection
	mutex         sync.RWMutex
}

type Subscription struct {
	active bool
	group  string
}

func (s *Subscription) IsActive() bool {
	return s.active
}

func (s *Subscription) GetGroup() string {
	return s.group
}

func NewConnectedClient(conn WebSocketConnection, id string) *ConnectedClient {
	return &ConnectedClient{
		id:            id,
		conn:          conn,
		mutex:         sync.RWMutex{},
		subscriptions: make(map[string]*Subscription),
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

	return client.subscriptions[topic]
}

func (client *ConnectedClient) SubscribeToTopic(topic string, group string) {
	client.mutex.Lock()
	defer client.mutex.Unlock()

	subscription := &Subscription{active: true}
	if len(group) > 0 {
		subscription.group = group
	}

	client.subscriptions[topic] = subscription
}

func (client *ConnectedClient) UnsubscribeFromTopic(topic string) {
	client.mutex.Lock()
	defer client.mutex.Unlock()

	delete(client.subscriptions, topic)
}
