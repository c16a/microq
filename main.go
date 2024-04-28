package main

import (
	"github.com/c16a/microq/broker"
	"github.com/c16a/microq/interfaces"
	"github.com/c16a/microq/storage"
	"log"
)

func main() {
	b := broker.NewBroker()

	var storageProvider = storage.NewFileStorageProvider("/tmp/microq")
	defer storageProvider.Close()

	go interfaces.RunWs(b, storageProvider)
	log.Fatal(interfaces.RunTcp(b, storageProvider))
}
