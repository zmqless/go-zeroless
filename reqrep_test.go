package zeroless

import "testing"

func TestPingPong(t *testing.T) {
	replier, _ := NewServer(1044).Rep()
	requester, _ := NewClient().ConnectLocal(1044).Req()

	pings := [][]string{
		[]string{"ping1"},
		[]string{"ping2"},
		[]string{"ping11", "ping12"},
		[]string{"ping21", "ping22"},
	}
	pongs := [][]string{
		[]string{"pong1"},
		[]string{"pong2"},
		[]string{"pong11", "ping12"},
		[]string{"pong21", "pong22"},
	}

	for i, ping := range pings {
		requester <- ping
		result := <-replier
		checkExchangedData(t, result, ping)
		replier <- pongs[i]
		result = <-requester
		checkExchangedData(t, result, pongs[i])
	}
}
