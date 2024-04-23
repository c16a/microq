package conn

import "github.com/gorilla/websocket"

type WebsocketConnection struct {
	conn *websocket.Conn
}

func NewWebsocketConnection(conn *websocket.Conn) *WebsocketConnection {
	return &WebsocketConnection{conn: conn}
}

func (wc *WebsocketConnection) WriteMessage(data []byte) error {
	return wc.conn.WriteMessage(websocket.TextMessage, data)
}
