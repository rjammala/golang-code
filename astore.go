package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync/atomic"
	"syscall"
	"time"
)

var version atomic.Value

func refreshVersion() {
	b, err := ioutil.ReadFile("temp.txt")
	if err != nil {
		log.Fatal(err)
	}
	version.Store(strings.TrimSpace(string(b)))
}

func init() {
	refreshVersion()
}

func main() {

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGHUP)
		for {
			select {
			case <-c:
				refreshVersion()
			}
		}
	}()

	timer := time.NewTicker(2 * time.Second)
	for {
		select {
		case <-timer.C:
			fmt.Printf("version: %s\n", version.Load().(string))
		}
	}

}
