package storage

import (
	"encoding/json"
	"errors"
	"github.com/c16a/microq/config"
	"github.com/c16a/microq/events"
	"io/fs"
	"os"
)

type FileProvider struct {
	rootDir string
}

func (f *FileProvider) SaveMessage(event *events.PubEvent) error {
	fSys := os.DirFS(f.rootDir)
	fInfo, err := fs.Stat(fSys, event.Topic)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			if dirCreateErr := os.Mkdir(f.rootDir+"/"+event.Topic, 0644); dirCreateErr != nil {
				return dirCreateErr
			}
		} else {
			return err
		}
	}
	if !fInfo.IsDir() {
		return errors.New("not a directory")
	} else {
		fullPath := f.rootDir + "/" + event.Topic
		f, err := os.Create(fullPath)
		if err != nil {
			return err
		}

		eventBytes, err := json.Marshal(event)
		if err != nil {
			return err
		}
		_, err = f.Write(eventBytes)
		return err
	}
}

func (f *FileProvider) Close() error {
	return nil
}

func NewFileStorageProvider(c *config.Config) *FileProvider {
	return &FileProvider{rootDir: c.Storage.RootDir}
}
