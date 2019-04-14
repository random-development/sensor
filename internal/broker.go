package internal

//go:generate mockgen -destination=../mocks/broker_mock.go -package=mocks github.com/random-development/sensor/internal Broker

// Broker is used to broadcast messages on particular topics between goroutines
type Broker interface {
	Pub(message interface{}, topic ...string)
	Sub(topic ...string) chan interface{}
}
