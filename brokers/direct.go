package brokers

import (
	"errors"
	"sync"
)

type DirectBroker struct {
	mu          sync.Mutex
	subscribers map[string]chan string
	closed      bool
	exited      chan bool
}

func NewDirectBroker() Broker {
	return &DirectBroker{
		subscribers: make(map[string]chan string),
		exited:      make(chan bool),
	}
}

func (b *DirectBroker) Subscribe(queuName string) (chan string, error) {
	if b.closed {
		return nil, errors.New("channel is already closed")
	}

	defer b.mu.Unlock()
	b.mu.Lock()

	_, exist := b.subscribers[queuName]
	if exist {
		return nil, errors.New("queue is already subscribed")
	}

	b.subscribers[queuName] = make(chan string)
	return b.subscribers[queuName], nil
}

func (b *DirectBroker) Publish(queuName string, message string) error {
	if b.closed {
		return errors.New("queue is already closed")
	}

	msgCh, exist := b.subscribers[queuName]
	if !exist {
		return errors.New("invalid queue name")
	}
	defer b.mu.Unlock()

	b.mu.Lock()
	msgCh <- message
	return nil
}

func (b *DirectBroker) Close() error {
	if b.closed {
		return errors.New("channel is already closed")
	}

	for key, _ := range b.subscribers {
		close(b.subscribers[key])
	}

	b.closed = true
	b.exited <- true
	return nil
}

func (b *DirectBroker) SubscribeExitedChannel() chan bool {
	return b.exited
}
