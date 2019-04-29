package internal

//go:generate mockgen -destination=../mocks/websocket_publisher_mock.go -package=mocks github.com/random-development/sensor/internal Conn,Dialer

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// WebSocketPublisher pushes Measurements via WebSockets
type WebSocketPublisher struct {
	conn Conn
}

// Dialer is used to perform WebSocket dial operation
type Dialer interface {
	Dial(url string, requestHeader http.Header) (Conn, *http.Response, error)
}

// Conn represents WebSocket connection
type Conn interface {
	WriteJSON(m interface{}) error
}

// DialerWrapper adapts gorilla.websocket.Dialer to internal interface
type DialerWrapper struct {
	Dialer *websocket.Dialer
}

// Dial wraps dialer implementation call so the returned Conn interaface is the internal one,
// not the websocket.Conn
func (d DialerWrapper) Dial(url string, requestHeader http.Header) (Conn, *http.Response, error) {
	return d.Dialer.Dial(url, requestHeader)
}

// MakeWebSocketPublisher builds WebSocketPublisher
func MakeWebSocketPublisher(url string, dialer Dialer) (WebSocketPublisher, error) {
	for {
		conn, _, err := dialer.Dial(url, nil)

		if err == nil {
			return WebSocketPublisher{conn}, nil
		} else {
			Log.Warnf("Couldn't connect to %s, %v, retrying in 5s", url, err)
			time.Sleep(5 * time.Second)
		}
	}
}

// Publish sends JSON message with Measurement via WebSocket
func (p WebSocketPublisher) Publish(m Measurement) error {
	Log.Debugf("Publishing via WebSocket: %s", m.String())
	if err := p.conn.WriteJSON(m); err != nil {
		Log.Errorf("Couldn't publish measurement: %v", err)
		return err
	}

	return nil
}
