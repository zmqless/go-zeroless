package zeroless

import "testing"

func TestDistribute(t *testing.T) {
	pusher, _ := NewServer(1034).Push()
	puller, _ := NewClient().ConnectLocal(1034).Pull()

	msgs := [][]string{
		[]string{"msg"},
		[]string{"msg11", "msg12"},
		[]string{"msg21", "msg22"},
	}

	for _, msg := range msgs {
		pusher <- msg
		result := <-puller
		checkExchangedData(t, result, msg)
	}
}
