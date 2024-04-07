package storage

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/c16a/microq/events"
	"github.com/dgraph-io/badger/v4"
	"github.com/dgraph-io/badger/v4/options"
	"log"
)

type BadgerProvider struct {
	db *badger.DB
}

func NewBadgerProvider() *BadgerProvider {
	var opts badger.Options
	opts = badger.DefaultOptions("/tmp/badger")
	opts.Compression = options.ZSTD
	opts.NumVersionsToKeep = 0
	opts.CompactL0OnClose = true
	opts.InMemory = false

	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}
	return &BadgerProvider{db: db}
}

func (b *BadgerProvider) SaveMessage(event *events.PubEvent) error {
	buf := new(bytes.Buffer)
	encoder := gob.NewEncoder(buf)
	err := encoder.Encode(event)
	if err != nil {
		return err
	}

	storageKey := fmt.Sprintf("%s:%s", event.Topic, event.PacketId)
	return b.db.Update(func(txn *badger.Txn) error {
		entry := badger.NewEntry([]byte(storageKey), buf.Bytes())
		return txn.SetEntry(entry)
	})
}

func (b *BadgerProvider) Close() error {
	return b.db.Close()
}
