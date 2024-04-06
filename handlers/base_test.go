package handlers

type TestingWebSocketConnection struct {
}

func (t *TestingWebSocketConnection) WriteMessage(messageType int, data []byte) error {
	return nil
}

func (t *TestingWebSocketConnection) ReadMessage() (messageType int, p []byte, err error) {
	return 0, nil, nil
}
