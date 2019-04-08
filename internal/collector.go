package internal

import (
	"fmt"
	"time"

	"github.com/cskr/pubsub"
)

// Collector represent a goroutine responsible for collecting
// particular resource's measurements
type Collector struct {
	probe Probe
	ps    *pubsub.PubSub
}

// MakeCollector builds collector with a given Probe
func MakeCollector(p Probe, ps *pubsub.PubSub) Collector {
	return Collector{
		probe: p,
		ps:    ps}
}

// Run starts a goroutine collecting measurements
func (c *Collector) Run(interval time.Duration) {
	go func() {
		fmt.Printf("Starting Collector with Probe(%s)\n", c.probe.Resource())
		for range time.Tick(interval) {
			m, err := c.probe.Measure()
			if err != nil {
				fmt.Printf("Failed to collect %s\n", c.probe.Resource())
			}
			c.ps.Pub(*m, c.probe.Resource())
		}
	}()
}
