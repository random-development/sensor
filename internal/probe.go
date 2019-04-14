package internal

//go:generate mockgen -destination=../mocks/probe_mock.go -package=mocks github.com/random-development/sensor/internal Probe

import (
	"time"
)

// Probe represents a single metric and a way in which it may be gathered
type Probe interface {
	MetricName() string
	Measure() (Measurement, error)
}

// RunProbe starts a goroutine collecting measurements
func RunProbe(probe Probe, broker Broker, interval time.Duration) {
	go func() {
		Log.Infof("Starting %s collector", probe.MetricName())
		for range time.Tick(interval) {
			m, err := probe.Measure()
			if err != nil {
				Log.Warnf("Failed to collect %s", probe.MetricName())
			}
			broker.Pub(m, probe.MetricName())
			Log.Debugf("Sent measurement to broker: %s", m.String())
		}
	}()
}
