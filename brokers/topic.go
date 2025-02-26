package brokers

import (
	"errors"
	"sync"
)

type TopicBroker struct {
	mu          sync.Mutex
	subscribers map[string]map[string]chan string
	closed      bool
	exited      chan bool
}

func NewTopicsBroker() *TopicBroker {
	return &TopicBroker{
		subscribers: make(map[string]map[string]chan string),
		exited:      make(chan bool),
	}
}

func (b *TopicBroker) Subscribe(queuName string, clientName string) (chan string, error) {
	if b.closed {
		return nil, errors.New("channel is already closed")
	}

	defer b.mu.Unlock()
	b.mu.Lock()

	subscribers, exist := b.subscribers[queuName]
	if exist {
		_, subscriberExist := subscribers[clientName]
		if subscriberExist {
			return nil, errors.New("client is already subscribed")
		}
	} else {
		b.subscribers[queuName] = make(map[string]chan string)
	}
	b.subscribers[queuName][clientName] = make(chan string)
	return b.subscribers[queuName][clientName], nil
}

func (b *TopicBroker) Publish(queuName string, message string) error {
	if b.closed {
		return errors.New("queue is already closed")
	}

	channelMap, exist := b.subscribers[queuName]
	if !exist {
		return errors.New("invalid queue name")
	}
	defer b.mu.Unlock()

	b.mu.Lock()

	for _, ch := range channelMap {
		ch <- message
	}
	return nil
}

func (b *TopicBroker) Close() error {
	if b.closed {
		return errors.New("channel is already closed")
	}

	for _, chMap := range b.subscribers {
		for _, ch := range chMap {
			close(ch)
		}
	}

	b.closed = true
	b.exited <- true
	return nil
}

func (b *TopicBroker) SubscribeExitedChannel() chan bool {
	return b.exited
}
