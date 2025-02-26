package brokers

type Broker interface {
	Publish(queuName string, message string) error
	Subscribe(queuName string) (chan string, error)
	SubscribeExitedChannel() chan bool
	Close() error
}
