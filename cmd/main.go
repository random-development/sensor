package main

import (
	"os"
	"os/signal"
	"time"

	"github.com/cskr/pubsub"
	"github.com/gorilla/websocket"
	sensor "github.com/random-development/sensor/internal"
)

func runProbes(config sensor.Config, broker sensor.Broker) {
	type probeFactory func() sensor.Probe

	var probeFactories = map[string]probeFactory{
		"memory": func() sensor.Probe { return sensor.MemProbe{} },
		"cpu":    func() sensor.Probe { return sensor.CPUProbe{} },
	}

	for _, probeConfig := range config.Probes {
		probeFactory, ok := probeFactories[probeConfig.Type]
		if !ok {
			sensor.Log.Errorf("Couldn't find factory for probe type: %s", probeConfig.Type)
			continue
		}
		sensor.RunProbe(probeFactory(), broker, time.Duration(probeConfig.Interval)*time.Second)
	}
}

func runPublishers(config sensor.Config, broker sensor.Broker, done chan bool) {
	type publisherFactory func(sensor.PublisherConfig) (sensor.Publisher, error)

	var publisherFactories = map[string]publisherFactory{
		"websocket": func(c sensor.PublisherConfig) (sensor.Publisher, error) {
			return sensor.MakeWebSocketPublisher(c.URL, sensor.DialerWrapper{Dialer: websocket.DefaultDialer})
		},
	}

	for _, publisherConfig := range config.Publishers {
		factory, ok := publisherFactories[publisherConfig.Type]
		if !ok {
			sensor.Log.Errorf("Couldn't find factory for publisher: %s", publisherConfig.Type)
			continue
		}

		for _, probeConfig := range config.Probes {
			publisher, err := factory(publisherConfig)
			if err != nil {
				sensor.Log.Errorf("Couldn't build publisher %s, %s", publisherConfig.Type, err)
				continue
			}
			sensor.RunPublisher(probeConfig.Type, publisher, broker, done)
		}
	}
}

func initInterruptHandler(done chan bool) {
	interruptCh := make(chan os.Signal, 1)
	signal.Notify(interruptCh, os.Interrupt)
	go func() {
		<-interruptCh
		sensor.Log.Info("Interrupt received, cleaning up")
		close(done)
	}()
}

func main() {
	config := sensor.ReadConfig()

	done := make(chan bool)
	initInterruptHandler(done)

	psBroker := pubsub.New(0)
	runProbes(config, psBroker)
	runPublishers(config, psBroker, done)

	<-done
}
