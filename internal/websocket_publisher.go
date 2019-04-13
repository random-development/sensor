package internal

//go:generate mockgen -destination=../mocks/websocket_publisher_mock.go -package=mocks github.com/random-development/sensor/internal Conn,Dialer

import (
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

type DialerWrapper struct {
	Dialer *websocket.Dialer
}

func (d DialerWrapper) Dial(url string, requestHeader http.Header) (Conn, *http.Response, error) {
	return d.Dialer.Dial(url, requestHeader)
}

// MakeWebSocketPublisher builds WebSocketPublisher
func MakeWebSocketPublisher(url string, dialer Dialer) (WebSocketPublisher, error) {
	conn, _, err := dialer.Dial(url, nil)
	if err != nil {
		log.Errorf("Error dialing to %s: %v", url, err)
		return WebSocketPublisher{}, err
	}

	return WebSocketPublisher{conn}, nil
}

// Publish sends JSON message with Measurement via WebSocket
func (p WebSocketPublisher) Publish(m Measurement) error {
	log.Debugf("Publishing via WebSocket: %s", m.String())
	if err := p.conn.WriteJSON(m); err != nil {
		log.Errorf("Couldn't publish measurement: %v", err)
		return err
	}

	return nil
}
