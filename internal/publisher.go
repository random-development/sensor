package internal

// Publisher receives measurements and publishes them to a specific sink
type Publisher interface {
	Run()
}
