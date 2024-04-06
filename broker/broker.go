package broker

import (
	"encoding/json"
	"github.com/c16a/microq/events"
	"math/rand"
	"sync"
)

type WebSocketConnection interface {
	WriteMessage(messageType int, data []byte) error
	ReadMessage() (messageType int, p []byte, err error)
}

type Broker struct {
	clients map[string]*ConnectedClient
	mutex   sync.RWMutex
}

func NewBroker() *Broker {
	return &Broker{
		clients: make(map[string]*ConnectedClient),
	}
}

func (broker *Broker) Connect(clientId string, client *ConnectedClient) {
	broker.mutex.Lock()
	defer broker.mutex.Unlock()

	broker.clients[clientId] = client
}

func (broker *Broker) Disconnect(clientId string) {
	broker.mutex.Lock()
	defer broker.mutex.Unlock()

	delete(broker.clients, clientId)
}

func (broker *Broker) Broadcast(event events.PubEvent) error {
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}

	var groupedClients = make(map[string][]*ConnectedClient, 0)
	for _, client := range broker.clients {
		subscription := client.GetEligibility(event.Topic)
		if subscription != nil {
			if subscription.group == "" {
				client.WriteDataMessage(data)
			} else {
				if _, ok := groupedClients[subscription.group]; !ok {
					groupedClients[subscription.group] = make([]*ConnectedClient, 0)
				}
				groupedClients[subscription.group] = append(groupedClients[subscription.group], client)
			}
		}
	}

	// Go to every group, and write the data to just one of them
	for _, clients := range groupedClients {
		idx := 0
		if len(clients) > 1 {
			idx = pickRandomNumber(0, len(clients)-1)
		}
		client := clients[idx]
		client.WriteDataMessage(data)
	}

	return nil
}

func pickRandomNumber(min int, max int) int {
	return min + rand.Intn(max-min)
}
