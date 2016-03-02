package main

import (
	"flag"
	"log"

	"github.com/patito/FileWatcher/configuration"
	"github.com/patito/FileWatcher/watcher"
)

func main() {

	configFlag := flag.String("config", "configuration/config.gcfg", "Configuration file")
	newConf, err := configuration.NewConfiguration(*configFlag)
	if err != nil {
		log.Fatal(err)
	}

	watcher, err := watcher.NewFileWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)

	if err = watcher.SetPath(newConf.Folder.Path); err != nil {
		log.Fatal(err)
	}
	go watcher.Listen()

	<-done
}
