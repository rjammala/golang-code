package main

import "log"
import "golang.org/x/exp/inotify"

var watchEventChannel chan *inotify.Event = make(chan *inotify.Event, 100)

func watchDir(name string) {
	watcher, err := inotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	err = watcher.AddWatch(name, inotify.IN_CREATE)
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case ev := <-watcher.Event:
			watchEventChannel <- ev
		case err := <-watcher.Error:
			log.Println("error:", err)
		}
	}
}

func main() {
	go watchDir("/tmp")
	go watchDir("/root")
	for {
		select {
		case ev := <-watchEventChannel:
			log.Println("event: ", ev)
		}
	}
}
