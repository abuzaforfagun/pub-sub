package brokers

import (
	"errors"
	"sync"
)

type DirectBroker[T any] struct {
	mu          sync.Mutex
	subscribers map[string]chan T
	closed      bool
	exited      chan bool
}

func NewDirectBroker[T any]() Broker[T] {
	return &DirectBroker[T]{
		subscribers: make(map[string]chan T),
		exited:      make(chan bool),
	}
}

func (b *DirectBroker[T]) Subscribe(queuName string) (chan T, error) {
	if b.closed {
		return nil, errors.New("channel is already closed")
	}

	defer b.mu.Unlock()
	b.mu.Lock()

	_, exist := b.subscribers[queuName]
	if exist {
		return nil, errors.New("queue is already subscribed")
	}

	b.subscribers[queuName] = make(chan T)
	return b.subscribers[queuName], nil
}

func (b *DirectBroker[T]) Publish(queuName string, message T) error {
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

func (b *DirectBroker[T]) Close() error {
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

func (b *DirectBroker[T]) SubscribeExitedChannel() chan bool {
	return b.exited
}
