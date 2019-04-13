package internal

//go:generate mockgen -destination=../mocks/websocket_publisher_mock.go -package=mocks github.com/random-development/sensor/internal Conn,Dialer

import (
	"fmt"
	"net/http"
)

// WebSocketPublisher pushes Measurements via WebSockets
type WebSocketPublisher struct {
	conn Conn
	ch   chan interface{}
}

type Dialer interface {
	Dial(url string, requestHeader http.Header) (Conn, *http.Response, error)
}

type Conn interface {
	WriteJSON(m interface{}) error
}

// MakeWebSocketPublisher builds WebSocketPublisher
func MakeWebSocketPublisher(metric, url string, dialer Dialer, ch chan interface{}) (WebSocketPublisher, error) {
	conn, _, err := dialer.Dial(url, nil)
	if err != nil {
		fmt.Printf("Error dialing to %s: %v", url, err)
		return WebSocketPublisher{}, err
	}

	return WebSocketPublisher{conn, ch}, nil
}

// Run starts goroutine which will receive measurements of one metric
// and publish them via WebSocket
func (p WebSocketPublisher) Run(done chan bool) {
	go func() {
		for {
			select {
			case m := <-p.ch:
				switch meas := m.(type) {
				case Measurement:
					p.handleMeasurement(meas)
				default:
					fmt.Println("Unknown type received")
				}
			case <-done:
				return
			}
		}
	}()
}

func (p WebSocketPublisher) handleMeasurement(m Measurement) {
	fmt.Printf("Handling measurement: %s\n", m.String())
	if err := p.conn.WriteJSON(m); err != nil {
		fmt.Printf("Couldn't publish measurement: %v\n", err)
	}
}

// Publish sends JSON message with Measurement via WebSocket
func (p WebSocketPublisher) publish(m Measurement) {
	fmt.Printf("Publishing measurement: %s\n", m.String())
}
