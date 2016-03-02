package watcher

import (
	"fmt"
	"net/http"

	"github.com/fsnotify/fsnotify"
)

type FileWatcher struct {
	*fsnotify.Watcher
}

func (watcher *FileWatcher) Listen() {
	for {
		select {
		case event := <-watcher.Events:
			if event.Op&fsnotify.Create == fsnotify.Create {
				http.Get("http://localhost:8000")
				fmt.Println("event:", event)
				fmt.Println("File Name:", event.Name)
			}

		case err := <-watcher.Errors:
			fmt.Println("error:", err)
		}
	}
}

func (watcher *FileWatcher) SetPath(path string) error {
	err := watcher.Add(path)
	if err != nil {
		return err
	}

	return nil
}

func NewFileWatcher() (*FileWatcher, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	return &FileWatcher{watcher}, nil
}
