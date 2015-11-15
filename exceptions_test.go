package zeroless

import "testing"

func TestPortBellowRange(t *testing.T) {
	port := 1023
	_, err := NewServer(port).Push()

	if err == nil {
		t.Error("Port", port, "is bellow 1024")
	}

	client := NewClient()
	client.ConnectLocal(port)
	_, err = client.Push()

	if err == nil {
		t.Error("Port", port, "is bellow 1024")
	}
}

func TestPortWithinRange(t *testing.T) {
	port := 1024
	_, err := NewServer(port).Push()

	if err != nil {
		t.Error("Port", port, "is within range")
	}

	client := NewClient()
	client.ConnectLocal(port)
	_, err = client.Push()

	if err != nil {
		t.Error("Port", port, "is within range")
	}
}

func TestPortAboveRange(t *testing.T) {
	port := 65536
	_, err := NewServer(port).Push()

	if err == nil {
		t.Error("Port", port, "is above 65535")
	}

	client := NewClient()
	client.ConnectLocal(port)
	_, err = client.Push()

	if err == nil {
		t.Error("Port", port, "is above 65535")
	}
}
