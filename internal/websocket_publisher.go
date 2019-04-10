package internal

//go:generate mockgen -destination=../mocks/websocket_publisher_mock.go -package=mocks github.com/random-development/sensor/internal Dialer,Conn

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// WebSocketPublisher pushes Measurements via WebSockets
type WebSocketPublisher struct {
	conn Conn
}

type Dialer interface {
	Dial(url string, requestHeader http.Header) (Conn, *http.Response, error)
}

type Conn interface {
	WriteJSON(m interface{}) error
}

type WebSocketDialer struct {
	dialer *websocket.Dialer
}

func (d WebSocketDialer) Dial(url string, requestHeader http.Header) (Conn, *http.Response, error) {
	return d.dialer.Dial(url, requestHeader)
}

type WebSocketConn struct {
	conn *websocket.Conn
}

// MakeWebSocketPublisher builds WebSocketPublisher
func MakeWebSocketPublisher(metric, url string, dialer Dialer) WebSocketPublisher {
	conn, _, err := dialer.Dial(url, nil)
	if err != nil {
		fmt.Printf("Error dialing to %s: %v", url, err)
	}

	return WebSocketPublisher{conn}
}

// Run starts goroutine which will receive measurements of one metric
// and publish them via WebSocket
func (p WebSocketPublisher) Run() {
	go func() {
		for {
		}
	}()
}

// Publish sends JSON message with Measurement via WebSocket
func (p WebSocketPublisher) publish(m Measurement) {
	fmt.Printf("Publishing measurement: %s\n", m.String())
}
