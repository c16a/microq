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
	fInfo, err := fs.Stat(fSys, "msg.txt")

	var offlineFile *os.File
	if err != nil {
		// Couldn't stat the file
		if errors.Is(err, fs.ErrNotExist) {
			// If file doesn't exist, create it.
			if offlineFile, err = os.Create(f.rootDir + "/msg.txt"); err != nil {
				return err
			}
		} else {
			// Any other error, throw
			return err
		}
	}
	if fInfo.IsDir() {
		return errors.New("found a directory instead of a file")
	}
	eventBytes, err := json.Marshal(event)
	if err != nil {
		return err
	}
	_, err = offlineFile.Write(eventBytes)
	return err
}

func (f *FileProvider) Close() error {
	return nil
}

func NewFileStorageProvider(c *config.Config) *FileProvider {
	return &FileProvider{rootDir: c.Storage.RootDir}
}
