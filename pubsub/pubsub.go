package pubsub

import (
	"sync"
)

type PubSub struct {
	mux         sync.RWMutex
	subscribers map[string][]chan string
}

func NewPubSub() *PubSub {
	return &PubSub{
		subscribers: make(map[string][]chan string),
	}
}

func (ps *PubSub) Subscribe(topic string) chan string {
	ps.mux.Lock()
	defer ps.mux.Unlock()

	// create a new channel with capacity 1
	ch := make(chan string, 1)

	// add new subscriber to the topic
	// any messages sent to the topic will be delivered to this new channel
	ps.subscribers[topic] = append(ps.subscribers[topic], ch)

	return ch
}

func (ps *PubSub) Publish(topic string, msg string) {
	ps.mux.RLock()
	defer ps.mux.RUnlock()

	// publish/send the message to the channel ch
	for _, ch := range ps.subscribers[topic] {
		select {
		case ch <- msg:
		default:
		}
	}
	// ch := make(chan string, 1)
	// ch <- msg
}
