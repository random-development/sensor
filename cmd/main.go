package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/cskr/pubsub"
	"github.com/random-development/sensor/internal"
)

func consumeMeasurements(ch chan interface{}) {
	for {
		event, ok := (<-ch).(internal.Measurement)
		if !ok {
			fmt.Printf("Wrong type received. Expected string\n")
		}

		fmt.Printf("Event received: %s\n", event.String())
	}
}

func main() {
	ps := pubsub.New(0)

	// spawn collectors and run their goroutines
	memCollector := internal.MakeCollector(internal.MemProbe{}, ps)
	memCollector.Run(time.Second)

	cpuCollector := internal.MakeCollector(internal.CPUProbe{}, ps)
	cpuCollector.Run(2 * time.Second)

	// spawn goroutines for consuming measurements generated by collectors
	go consumeMeasurements(ps.Sub(internal.MemProbe{}.Resource()))
	go consumeMeasurements(ps.Sub(internal.CPUProbe{}.Resource()))

	// handle interrupt
	interruptCh := make(chan os.Signal, 1)
	signal.Notify(interruptCh, os.Interrupt)
	done := make(chan bool)
	go func() {
		<-interruptCh
		fmt.Println("Interrupt received, cleaning up")
		close(done)
	}()
	<-done
}