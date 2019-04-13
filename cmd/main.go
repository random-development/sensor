package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/cskr/pubsub"
	"github.com/gorilla/websocket"
	"github.com/random-development/sensor/internal"
)

func main() {
	psBroker := pubsub.New(0)
	done := make(chan bool)

	// spawn collectors and run their goroutines
	memCollector := internal.MakeCollector(internal.MemProbe{}, psBroker)
	memCollector.Run(time.Second)

	// spawn goroutines for consuming measurements generated by collectors
	wsPublisher, _ := internal.MakeWebSocketPublisher("ws://demos.kaazing.com/echo",
		internal.DialerWrapper{Dialer: websocket.DefaultDialer})
	internal.RunPublisher("memory", wsPublisher, psBroker, done)

	// handle interrupt
	interruptCh := make(chan os.Signal, 1)
	signal.Notify(interruptCh, os.Interrupt)
	go func() {
		<-interruptCh
		fmt.Println("Interrupt received, cleaning up")
		done <- true
	}()
	<-done

	// TODO: wait for signals from all workers
	t := time.NewTimer(2 * time.Second)
	<-t.C
}
