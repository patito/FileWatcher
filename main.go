package main

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
)

type FileWatcher struct {
	*fsnotify.Watcher
}

func (watcher *FileWatcher) FileWatcherEvents() {
	for {
		select {
		case event := <-watcher.Events:
			fmt.Println("event:", event)
			fmt.Println("File Name:", event.Name)
		case err := <-watcher.Errors:
			fmt.Println("error:", err)
		}
	}
}

func NewFileWatcher() *FileWatcher {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	fileWatcher := &FileWatcher{watcher}

	return fileWatcher
}

func main() {
	watcher := NewFileWatcher()
	defer watcher.Close()

	done := make(chan bool)
	go watcher.FileWatcherEvents()

	err := watcher.Add("/tmp/foo")
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
