package main

import (
	"github.com/c16a/microq/broker"
	"github.com/c16a/microq/handlers"
	"github.com/c16a/microq/storage"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func main() {
	var storageProvider = storage.NewBadgerProvider()
	defer storageProvider.Close()

	b := broker.NewBroker()
	var upgrader = websocket.Upgrader{}
	http.HandleFunc("/echo", echo(upgrader, b, storageProvider))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func echo(upgrader websocket.Upgrader, b *broker.Broker, sp storage.Provider) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		c, err := upgrader.Upgrade(writer, request, nil)
		if err != nil {
			return
		}
		defer c.Close()

		clientId := request.Header.Get("Client-Id")
		client := broker.NewConnectedClient(c, clientId)
		b.Connect(clientId, client)

		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseAbnormalClosure) {
					b.Disconnect(clientId)
					break
				} else {
					continue
				}
			}
			err = handlers.HandleMessage(client, b, sp, message)
			if err != nil {
				continue
			}
		}
	}
}
