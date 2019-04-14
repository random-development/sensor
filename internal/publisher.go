package internal

//go:generate mockgen -destination=../mocks/publisher_mock.go -package=mocks github.com/random-development/sensor/internal Publisher

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
func RunPublisher(topic string, p Publisher, b Broker, done chan bool) {
	measCh := b.Sub(topic)
	l := Log.WithFields(logrus.Fields{
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
