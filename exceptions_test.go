package zeroless

import "testing"

func TestPortUnderRange(t *testing.T) {
	port := 1023
	_, err := NewServer(port).Push()

	if err == nil {
		t.Error("Port", port, "is under 1024")
	}
}

func TestPortOnRange(t *testing.T) {
	port := 1024
	_, err := NewServer(port).Push()

	if err != nil {
		t.Error("Port", port, "is not on range")
	}
}

func TestPortAfterRange(t *testing.T) {
	port := 65536
	_, err := NewServer(port).Push()

	if err == nil {
		t.Error("Port", port, "is after 65535")
	}
}
