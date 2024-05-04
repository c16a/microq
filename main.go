package main

import (
	"encoding/json"
	"flag"
	"github.com/c16a/microq/broker"
	"github.com/c16a/microq/config"
	"github.com/c16a/microq/interfaces"
	"github.com/c16a/microq/storage"
	"log"
	"os"
)

var configFileName string

func init() {
	flag.StringVar(&configFileName, "config", "config.json", "configuration file")
}

func main() {
	flag.Parse()

	c, err := ParseJsonFile[config.Config](configFileName)
	if err != nil {
		panic(err)
	}

	b := broker.NewBroker()

	var storageProvider = storage.NewFileStorageProvider(c)
	defer storageProvider.Close()

	go interfaces.RunWs(b, storageProvider)
	log.Fatal(interfaces.RunTcp(b, storageProvider))
}

func ParseJsonFile[T any](path string) (*T, error) {
	var config T
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
