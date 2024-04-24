package conn

import "net"

type TCPConnection struct {
	conn *net.TCPConn
}

func NewTCPConnection(conn *net.TCPConn) *TCPConnection {
	return &TCPConnection{conn: conn}
}

func (tc *TCPConnection) WriteMessage(data []byte) error {
	_, err := tc.conn.Write(data)
	return err
}
