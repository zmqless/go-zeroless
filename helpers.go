package zeroless

import (
	"strconv"
)

func connectZmqSock(sock socket, ip string, port int) {
	if sock == nil {
		panic("Sock is nil for connecting")
	}

	endpoint := "tcp://" + ip + ":" + strconv.Itoa(port)
	sock.Connect(endpoint)
}

func disconnectZmqSock(sock socket, ip string, port int) {
	if sock == nil {
		panic("Sock is nil for disconnecting")
	}

	endpoint := "tcp://" + ip + ":" + strconv.Itoa(port)
	sock.Disconnect(endpoint)
}

func bindZmqSock(sock socket, port int) {
	if sock == nil {
		panic("Sock is nil for binding")
	}

	endpoint := "tcp://*:" + strconv.Itoa(port)
	sock.Bind(endpoint)
}

func unbindZmqSock(sock socket, port int) {
	if sock == nil {
		panic("Sock is nil for unbinding")
	}

	endpoint := "tcp://*:" + strconv.Itoa(port)
	sock.Unbind(endpoint)
}
