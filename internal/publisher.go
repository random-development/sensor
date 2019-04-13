package internal

import (
	"reflect"

	"github.com/sirupsen/logrus"
)

// Publisher receives measurements and publishes them somewhere
type Publisher interface {
	Publish(Measurement) error
}

// RunPublisher starts goroutine which will receive measurements of one metric
// and publish them
func RunPublisher(topic string, p Publisher, b Broker, done chan struct{}) {
	measCh := b.Sub(topic)
	l := log.WithFields(logrus.Fields{
		"publisher": reflect.TypeOf(p).Name(),
		"metric":    topic,
	})
	go func() {
		l.Info("Started")
		for {
			select {
			case m := <-measCh:
				switch meas := m.(type) {
				case Measurement:
					p.Publish(meas)
				default:
					l.Debug("Unknown type received")
				}
			case <-done:
				l.Info("Done")
				return
			}
		}
	}()
}
