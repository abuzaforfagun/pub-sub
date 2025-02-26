package main

import (
	"fmt"
	"sync"

	"github.com/abuzaforfagun/pub-sub/brokers"
)

func main() {
	directBroker := brokers.NewDirectBroker()

	musicListener, err := directBroker.Subscribe("music")
	if err != nil {
		fmt.Println(err)
	}
	sportListener, err := directBroker.Subscribe("sports")
	if err != nil {
		fmt.Println(err)
	}

	go func() {
		for val := range musicListener {
			fmt.Println(val)
		}
	}()

	go func() {
		for val := range sportListener {
			fmt.Println(val)
		}
	}()

	var wg sync.WaitGroup

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
