package interfaces

import (
	"github.com/c16a/microq/broker"
	"github.com/c16a/microq/conn"
	"github.com/c16a/microq/handlers"
	"github.com/c16a/microq/storage"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func RunWs(b *broker.Broker, storageProvider storage.Provider) {
	var upgrader = websocket.Upgrader{}
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		defer c.Close()

		wsConn := conn.NewWebsocketConnection(c)
		client := broker.NewUnidentifiedClient(wsConn)

		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseAbnormalClosure) {
					b.Disconnect(client)
					break
				} else {
					continue
				}
			}
			err = handlers.HandleMessage(client, b, storageProvider, message)
			if err != nil {
				continue
			}
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
