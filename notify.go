package main

import (
	"golang.org/x/exp/inotify"
	"log"
	"syscall"
	"os"
	"os/exec"
	"fmt"
)


var targetDir = "/home/tarr/dev/go/diffwatcher/gitcabinet_test_repo"

func main(){
		watcher, err := inotify.NewWatcher()
		if err != nil {
			log.Fatal(err)
		}

		err = watcher.AddWatch(targetDir, syscall.IN_ALL_EVENTS)

		if err != nil {
			log.Fatal(err)
		}

		for {
			select {
			case ev := <-watcher.Event:
				log.Println("event:", ev)
				GitPush()
			case err := <-watcher.Error:
				log.Println("error:", err)
			}
		}
}

func GitPush() {
	// change directory to target directory.  

/*
	out, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()	
	if err != nil {
    	fmt.Println(err)
        os.Exit(1)
    }
	target_path := string(out)
*/

	os.Chdir(targetDir)

	out, err := exec.Command("git", "push", "origin", "master").Output()	

	if err != nil {
    	fmt.Println(err)
        os.Exit(1)
    }

    fmt.Println(string(out))
}
