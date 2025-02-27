package main

import (
	"fmt"
	"sync"

	"github.com/abuzaforfagun/pub-sub/brokers"
)

func main() {
	// directBrokerExample()
	topicBrokkerExample()
}

func topicBrokkerExample() {
	topicBroker := brokers.NewTopicsBroker[string]()
	musicListener1, err := topicBroker.Subscribe("music", "1")
	if err != nil {
		fmt.Println(err)
	}
	musicListener2, err := topicBroker.Subscribe("music", "2")
	if err != nil {
		fmt.Println(err)
	}
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for msg := range musicListener1 {
			defer wg.Done()
			wg.Add(1)
			fmt.Println("Listener 1: ", msg)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for msg := range musicListener2 {
			defer wg.Done()
			wg.Add(1)
			fmt.Println("Listener 2: ", msg)
		}
	}()

	wg.Add(1)
	go func() {
		for {
			select {
			case <-topicBroker.SubscribeExitedChannel():
				fmt.Println("channel exited")
				wg.Done()
				return
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		topicBroker.Publish("music", "Going to release new music")
		topicBroker.Publish("music", "Released new music")
		topicBroker.Close()
	}()

	wg.Wait()
}

func directBrokerExample() {
	directBroker := brokers.NewDirectBroker[string]()

	musicListener, err := directBroker.Subscribe("music")
	if err != nil {
		fmt.Println(err)
	}
	sportListener, err := directBroker.Subscribe("sports")
	if err != nil {
		fmt.Println(err)
	}
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range musicListener {
			defer wg.Done()
			wg.Add(1)
			fmt.Println(val)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range sportListener {
			defer wg.Done()
			wg.Add(1)
			fmt.Println(val)
		}
	}()

	wg.Add(1)
	go func() {
		for {
			select {
			case <-directBroker.SubscribeExitedChannel():
				wg.Done()
				return
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		directBroker.Publish("music", "New song of Lynkin Park is just released")
		directBroker.Publish("music", "New song of Justin Biber is just released")
		directBroker.Publish("news", "New song of Justin Biber is just released")
		directBroker.Publish("sports", "Bangladesh won the match")
		directBroker.Close()
	}()

	wg.Wait()
}
