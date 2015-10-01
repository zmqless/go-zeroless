package zeroless

import (
	"testing"
	"time"
)

func TestPublish(t *testing.T) {
	publisher, _ := NewServer(1025).Pub()
	subscriber, _ := NewClient().ConnectLocal(1025).Sub()

	msgs := [][]string{
		[]string{"msg"},
		[]string{"msg11", "msg12"},
		[]string{"msg21", "msg22"},
	}

	time.Sleep(10 * time.Millisecond)
	for _, msg := range msgs {
		publisher <- msg
		result := <-subscriber
		checkExchangedData(t, result, msg)
	}
}

func TestPublishWithTopic(t *testing.T) {
	publisherWithTopic, _ := NewServer(1026).PubWithTopic("urgent", true)
	subscriberWithTopics, _ := NewClient().ConnectLocal(1026).SubWithTopics([]string{"urgent"})

	msgs := [][]string{
		[]string{"msg"},
		[]string{"msg11", "msg12"},
		[]string{"msg21", "msg22"},
	}

	time.Sleep(10 * time.Millisecond)
	for _, msg := range msgs {
		publisherWithTopic <- msg
		result := <-subscriberWithTopics
		checkExchangedData(t, result, append([]string{"urgent"}, msg...))
	}
}

func TestPublishWithManyTopics(t *testing.T) {
	publisherWithoutTopic, _ := NewServer(1027).Pub()
	subscriberWithManyTopics, _ := NewClient().ConnectLocal(1027).SubWithTopics([]string{"a", "c"})

	msgs := [][]string{
		[]string{"a", "msg"},
		[]string{"b", "msg11", "msg12"},
		[]string{"c", "msg21", "msg22"},
	}

	time.Sleep(10 * time.Millisecond)
	for _, msg := range msgs {
		publisherWithoutTopic <- msg

		if msg[0] != "b" {
			result := <-subscriberWithManyTopics
			checkExchangedData(t, result, msg)
		}
	}
}
