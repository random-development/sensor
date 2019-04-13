package internal

import (
	"time"
)

// Collector represent a goroutine responsible for collecting
// particular resource's measurements
type Collector struct {
	probe  Probe
	broker Broker
}

// MakeCollector builds collector with a given Probe
func MakeCollector(p Probe, broker Broker) Collector {
	return Collector{
		probe:  p,
		broker: broker}
}

// Run starts a goroutine collecting measurements
func (c *Collector) Run(interval time.Duration) {
	go func() {
		log.Infof("Starting %s collector", c.probe.Resource())
		for range time.Tick(interval) {
			m, err := c.probe.Measure()
			if err != nil {
				log.Warnf("Failed to collect %s", c.probe.Resource())
			}
			c.broker.Pub(m, c.probe.Resource())
			log.Debugf("Sent measurement to broker: %s", m.String())
		}
	}()
}
