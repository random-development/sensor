package internal

import "fmt"

// Publisher receives measurements and publishes them to a specific sink
type Publisher interface {
	Publish(Measurement) error
}

// RunPublisher starts goroutine which will receive measurements of one metric
// and publish them
func RunPublisher(topic string, p Publisher, b Broker, done chan struct{}) {
	measCh := b.Sub(topic)
	go func() {
		for {
			select {
			case m := <-measCh:
				switch meas := m.(type) {
				case Measurement:
					p.Publish(meas)
				default:
					fmt.Println("Unknown type received")
				}
			case <-done:
				fmt.Println("Done")
				return
			}
		}
	}()
}
