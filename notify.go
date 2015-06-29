package main

import (
	"golang.org/x/exp/inotify"
	"log"
	"syscall"
)

func main(){
		watcher, err := inotify.NewWatcher()
		if err != nil {
			log.Fatal(err)
		}

		err = watcher.AddWatch("/tmp", syscall.IN_CREATE)

		if err != nil {
			log.Fatal(err)
		}

		for {
			select {
			case ev := <-watcher.Event:
				log.Println("event:", ev)
				log.Println("Name:", ev.Mask)
			case err := <-watcher.Error:
				log.Println("error:", err)
			}
		}
}
