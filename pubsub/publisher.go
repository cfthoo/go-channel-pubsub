package pubsub

func Publisher(ps *PubSub, topic string, msg string) {
	ps.Publish(topic, msg)
}
