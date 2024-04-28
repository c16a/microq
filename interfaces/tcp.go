package interfaces

import (
	"bufio"
	"github.com/c16a/microq/broker"
	"github.com/c16a/microq/conn"
	"github.com/c16a/microq/handlers"
	"github.com/c16a/microq/storage"
	"net"
)

func RunTcp(b *broker.Broker, storageProvider storage.Provider) error {
	listener, err := net.ListenTCP("tcp", &net.TCPAddr{Port: 8081})
	if err != nil {
		return err
	}

	for {
		c, err := listener.AcceptTCP()
		if err != nil {
			continue
		}

		scanner := bufio.NewScanner(c)

		tcpConn := conn.NewTCPConnection(c)
		client := broker.NewUnidentifiedClient(tcpConn)

		for scanner.Scan() {
			line := scanner.Text()

			err = handlers.HandleMessage(client, b, storageProvider, []byte(line))
			if err != nil {
				continue
			}
		}
	}
}
