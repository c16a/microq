package broker

import (
	"encoding/json"
	"sync"
)

type ConnectedClient struct {
	id              string
	subscriptions   []*Subscription
	unsubscriptions []*Subscription
	conn            GenericConnection
	mutex           sync.RWMutex
}

func NewConnectedClient(conn GenericConnection, id string) *ConnectedClient {
	return &ConnectedClient{
		id:              id,
		conn:            conn,
		mutex:           sync.RWMutex{},
		subscriptions:   make([]*Subscription, 0),
		unsubscriptions: make([]*Subscription, 0),
	}
}

func NewUnidentifiedClient(conn GenericConnection) *ConnectedClient {
	return NewConnectedClient(conn, "")
}

func (client *ConnectedClient) IsIdentified() bool {
	return client.id != ""
}

func (client *ConnectedClient) SetId(id string) {
	client.id = id
}

func (client *ConnectedClient) GetId() string {
	return client.id
}

func (client *ConnectedClient) WriteDataMessage(data []byte) error {
	client.mutex.Lock()
	defer client.mutex.Unlock()

	return client.conn.WriteMessage(data)
}

func (client *ConnectedClient) WriteInterface(v any) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return client.WriteDataMessage(data)
}

func (client *ConnectedClient) GetEligibility(topic string) *Subscription {
	client.mutex.RLock()
	defer client.mutex.RUnlock()

	// If topic matches any unsubscribed pattern,
	// it's not eligible
	for _, unsub := range client.unsubscriptions {
		if unsub.Matches(topic) {
			return nil
		}
	}

	// If topic is not unsubscribed as a part of any pattern,
	// check if it matches the subscribed ones
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

	// If this exact pattern has been previously unsubscribed from, remove that entry
	for i, unsub := range client.unsubscriptions {
		if unsub.pattern == pattern {
			client.unsubscriptions = append(client.unsubscriptions[:i], client.unsubscriptions[i+1:]...)
		}
	}
}

func (client *ConnectedClient) UnsubscribeFromPattern(pattern string) {
	client.mutex.Lock()
	defer client.mutex.Unlock()

	unsubscription := &Subscription{active: true, pattern: pattern}
	client.unsubscriptions = append(client.unsubscriptions, unsubscription)

	// If this exact pattern has been previously subscribed to, remove that entry
	for i, sub := range client.subscriptions {
		if sub.pattern == pattern {
			client.subscriptions = append(client.subscriptions[:i], client.subscriptions[i+1:]...)
		}
	}
}
