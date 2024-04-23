package main

import (
	"bufio"
	"github.com/c16a/microq/broker"
	"github.com/c16a/microq/conn"
	"github.com/c16a/microq/handlers"
	"github.com/c16a/microq/storage"
	"github.com/gorilla/websocket"
	"log"
	"net"
	"net/http"
)

func main() {
	var storageProvider = storage.NewBadgerProvider()
	defer storageProvider.Close()

	b := broker.NewBroker()
	go runWebSocketInterface(b, storageProvider)
	log.Fatal(runTcpInterface(b, storageProvider))
}

func runTcpInterface(b *broker.Broker, storageProvider storage.Provider) error {
	listener, err := net.ListenTCP("tcp", &net.TCPAddr{Port: 8081})
	if err != nil {
		return err
	}

	for {
		// Handles one TCP client
		c, err := listener.AcceptTCP()
		if err != nil {
			continue
		}

		scanner := bufio.NewScanner(c)

		for scanner.Scan() {
			line := scanner.Text()

			tcpConn := conn.NewTCPConnection(c)
			client := broker.NewConnectedClient(tcpConn, "unlabeled")

			err = handlers.HandleMessage(client, b, storageProvider, []byte(line))
			if err != nil {
				continue
			}
		}
	}
}

func runWebSocketInterface(b *broker.Broker, storageProvider storage.Provider) {
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

		wsConn := conn.NewWebsocketConnection(c)
		client := broker.NewConnectedClient(wsConn, clientId)
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
