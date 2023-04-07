package main

import (
	"fmt"
	"sync"

	"go-channel/pubsub"
)

func main() {

	// initiate a new pubsub
	ps := pubsub.NewPubSub()

	// subscribe to the topics
	ch1 := ps.Subscribe("gaming")
	ch2 := ps.Subscribe("education")

	var wg sync.WaitGroup

	wg.Add(2)

	// start to listen for incoming message
	// set the wg to done after receiving the message
	go pubsub.Subscriber(ch1, &wg)
	go pubsub.Subscriber(ch2, &wg)

	// publish messages to the topics gaming/education that we subscribed earlier
	pubsub.Publisher(ps, "gaming", "Hello CF, we have a new game released today!")
	pubsub.Publisher(ps, "education", "Hello CF, welcome to the go-channel course!")

	// Close the channel and exit the goroutines
	// u can remove this based on your need
	close(ch1)
	close(ch2)

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("Program finished...")
}
