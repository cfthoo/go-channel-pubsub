package pubsub

import (
	"fmt"
	"sync"
)

func Subscriber(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case msg, ok := <-ch:
			if !ok {
				// something wrong...
				return
			}
			fmt.Printf("Received message on channel: %s\n", msg)
		default:
		}
	}
}
