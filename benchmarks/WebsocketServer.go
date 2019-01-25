package benchmarks

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		if "ws://"+r.Host == r.Header.Get("Origin") {
			return true
		}
		return false
	},
}

type websocketsServer interface {
	SendAndReceiveMessage()
	SetMessageSize(messageSize)
	CloseConnection()
	WaitForInitializaion()
}

type messageSize uint8

const (
	small messageSize = iota
	medium
	big
)
