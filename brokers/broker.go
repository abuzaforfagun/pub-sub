package brokers

type Broker[T any] interface {
	Publish(queuName string, message T) error
	Subscribe(queuName string) (chan T, error)
	SubscribeExitedChannel() chan bool
	Close() error
}
