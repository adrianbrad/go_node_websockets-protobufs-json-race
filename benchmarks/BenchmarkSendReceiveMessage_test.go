package benchmarks

import (
	"os/exec"
	"testing"
)

func BenchmarkSendAndReceiveProto(b *testing.B) {
	b.StopTimer()

	teardownTestCase, ws := setupTestCase("Proto")
	defer teardownTestCase()

	b.StartTimer()

	b.Run("Send and Receive Small Proto", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			ws.SendAndReceiveMessage()
		}
	})

	ws.SetMessageSize(medium)

	b.Run("Send and Receive Medium Proto", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			ws.SendAndReceiveMessage()
		}
	})

	ws.SetMessageSize(big)

	b.Run("Send and Receive Big Proto", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			ws.SendAndReceiveMessage()
		}
	})
}

func BenchmarkSendAndReceiveJson(b *testing.B) {
	b.StopTimer()

	teardownTestCase, ws := setupTestCase("Json")
	defer teardownTestCase()

	b.StartTimer()

	b.Run("Send and Receive Small Json", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			ws.SendAndReceiveMessage()
		}
	})

	ws.SetMessageSize(medium)

	b.Run("Send and Receive Medium Json", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			ws.SendAndReceiveMessage()
		}
	})

	ws.SetMessageSize(big)

	b.Run("Send and Receive Big Json", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			ws.SendAndReceiveMessage()
		}
	})
}

func setupTestCase(client string) (func(), websocketsServer) {
	cmd := exec.Command("node", "../ws-client/client"+client+".js")
	var ws websocketsServer

	switch client {
	case "Json":
		ws = NewJson()
	case "Proto":
		ws = NewProto()
	}

	ws.SetMessageSize(small)

	go cmd.Run()

	ws.WaitForInitializaion() //this waits for connection initialization
	return func() {
		cmd.Process.Kill()
		ws.CloseConnection()
	}, ws
}
