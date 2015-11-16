package zeroless

import (
	"errors"
	"strconv"
)

func checkPortInRange(port int) error {
	if port < 1024 || port > 65535 {
		return errors.New("Port " + strconv.Itoa(port) + " is invalid, choose one between 1024 and 65535")
	}

	return nil
}

func connectZmqSock(sock socket, ip string, port int) error {
	if sock == nil {
		panic("Sock is nil for connecting")
	}

	err := checkPortInRange(port)
	if err != nil {
		return err
	}

	endpoint := "tcp://" + ip + ":" + strconv.Itoa(port)
	sock.Connect(endpoint)

	return nil
}

func disconnectZmqSock(sock socket, ip string, port int) error {
	if sock == nil {
		panic("Sock is nil for disconnecting")
	}

	err := checkPortInRange(port)
	if err != nil {
		return err
	}

	endpoint := "tcp://" + ip + ":" + strconv.Itoa(port)
	sock.Disconnect(endpoint)

	return nil
}

func bindZmqSock(sock socket, port int) error {
	if sock == nil {
		panic("Sock is nil for binding")
	}

	err := checkPortInRange(port)
	if err != nil {
		return err
	}

	endpoint := "tcp://*:" + strconv.Itoa(port)
	sock.Bind(endpoint)

	return nil
}

func unbindZmqSock(sock socket, port int) error {
	if sock == nil {
		panic("Sock is nil for unbinding")
	}

	err := checkPortInRange(port)
	if err != nil {
		return err
	}

	endpoint := "tcp://*:" + strconv.Itoa(port)
	sock.Unbind(endpoint)

	return nil
}
