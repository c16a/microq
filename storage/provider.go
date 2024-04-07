package storage

import "github.com/c16a/microq/events"

type Provider interface {
	SaveMessage(event *events.PubEvent) error
	Close() error
}
