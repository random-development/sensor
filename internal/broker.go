package internal

//go:generate mockgen -destination=../mocks/broker_mock.go -package=mocks github.com/random-development/sensor/internal Broker

type Broker interface {
	Pub(message interface{}, topic ...string)
	Sub(topic ...string) chan interface{}
}
