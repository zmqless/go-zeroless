package zeroless

import (
	zmq "github.com/pebbe/zmq4"
)

type socket interface {
	Bind(string) error
	Unbind(string) error
	Connect(string) error
	Disconnect(string) error
}

type sockInitialiser interface {
	init(socket) error
}

type sock struct {
	initialiser sockInitialiser
	ctx         *zmq.Context
	zmqSock     *zmq.Socket
}

func (this *sock) setInit(initialiser sockInitialiser) {
	this.initialiser = initialiser
}

func (this sock) getZmqSock() *zmq.Socket {
	if this.ctx == nil {
		panic("Context is nil")
	}

	if this.zmqSock == nil {
		panic("Sock is nil")
	}

	return this.zmqSock
}

func (this sock) Ready() bool {
	return (this.ctx != nil && this.zmqSock != nil)
}

func (this *sock) setSock(pattern zmq.Type) error {
	var err error
	this.ctx, err = zmq.NewContext()
	if err != nil {
		panic(err)
	}

	this.zmqSock, err = zmq.NewSocket(pattern)
	if err != nil {
		panic(err)
	}

	return this.initialiser.init(this.zmqSock)
}

func (this sock) readFromZmqSock() []string {
	frames, err := this.getZmqSock().RecvMessage(0)
	if err != nil {
		panic(err)
	}
	return frames
}

func (this sock) writeToZmqSock(frames []string) {
	_, err := this.getZmqSock().SendMessage(frames)
	if err != nil {
		panic(err)
	}
}

func (this sock) getSender() chan []string {
	c := make(chan []string)

	go func() {
		defer close(c)

		for {
			frames := <-c
			this.writeToZmqSock(frames)
		}
	}()

	return c
}

func (this sock) getReceiver() chan []string {
	c := make(chan []string)

	go func() {
		defer close(c)

		for {
			c <- this.readFromZmqSock()
		}
	}()

	return c
}

func (this sock) getSenderReceiver() chan []string {
	c := make(chan []string)

	go func() {
		defer close(c)

		sender := this.getSender()
		receiver := this.getReceiver()

		for {
			select {
			case msg := <-c:
				sender <- msg
			case msg := <-receiver:
				c <- msg
			}
		}
	}()

	return c
}

func (this sock) getRequester() chan []string {
	c := make(chan []string)

	go func() {
		defer close(c)

		for {
			framesToSend := <-c
			this.writeToZmqSock(framesToSend)
			c <- this.readFromZmqSock()
		}
	}()

	return c
}
func (this sock) getReplier() chan []string {
	c := make(chan []string)

	go func() {
		defer close(c)

		for {
			c <- this.readFromZmqSock()
			framesToSend := <-c
			this.writeToZmqSock(framesToSend)
		}
	}()

	return c
}

func (this *sock) Pub() (chan<- []string, error) {
	err := this.setSock(zmq.PUB)

	if err != nil {
		return nil, err
	}

	return this.getSender(), nil
}

func (this *sock) PubWithTopic(topic string, embedTopic bool) (chan<- []string, error) {
	publisher, err := this.Pub()

	if err != nil {
		return nil, err
	}

	c := make(chan []string)

	go func() {
		defer close(c)

		for {
			msg := <-c

			if embedTopic {
				publisher <- append([]string{topic}, msg...)
			} else {
				topicLen := len(topic)
				if msg[0][:topicLen] != topic {
					errString := "If embedTopic argument is not set, then the " +
						"topic must be at the beginning of the first " +
						"part (i.e. frame) of every published message"
					panic(errString)
				}

				publisher <- msg
			}
		}
	}()

	return c, nil
}

func (this *sock) Sub() (<-chan []string, error) {
	err := this.setSock(zmq.SUB)

	if err != nil {
		return nil, err
	}

	err = this.zmqSock.SetSubscribe("")

	if err != nil {
		panic(err)
	}

	return this.getReceiver(), nil
}

func (this *sock) SubWithTopics(topics []string) (<-chan []string, error) {
	err := this.setSock(zmq.SUB)

	if err != nil {
		return nil, err
	}

	for _, topic := range topics {
		if topic == "" {
			panic("You cannot set an empty string to be a topic.")
		}

		err := this.zmqSock.SetSubscribe(topic)

		if err != nil {
			return nil, err
		}
	}

	return this.getReceiver(), nil
}

func (this *sock) Push() (chan<- []string, error) {
	err := this.setSock(zmq.PUSH)

	if err != nil {
		return nil, err
	}

	return this.getSender(), nil
}

func (this *sock) Pull() (<-chan []string, error) {
	err := this.setSock(zmq.PULL)

	if err != nil {
		return nil, err
	}

	return this.getReceiver(), nil
}

func (this *sock) Req() (chan []string, error) {
	err := this.setSock(zmq.REQ)

	if err != nil {
		return nil, err
	}

	return this.getRequester(), nil
}

func (this *sock) Rep() (chan []string, error) {
	err := this.setSock(zmq.REP)

	if err != nil {
		return nil, err
	}

	return this.getReplier(), nil
}

func (this *sock) Pair() (chan []string, error) {
	err := this.setSock(zmq.PAIR)

	if err != nil {
		return nil, err
	}

	return this.getSenderReceiver(), nil
}
