package main

import (
	"fmt"
	"math"
	"net/http"
	"os"
	"strconv"
	"time"

	"websockets-protobufs-json-race/proto"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
)

var (
	times = 20000
	count = 0
	data  []byte

	small_protoMessage = &message.Message{
		Integer:  9,
		Floating: 1.1,
	}

	medium_protoMessage = &message.Message{
		Integer:  math.MaxInt32,
		Floating: math.MaxFloat32,
		Pairs: []*message.TextNumberPair{
			{
				Text:   "Lorem ipsum",
				Number: math.MaxInt32,
			},
		},
	}

	big_protoMesage = &message.Message{
		Integer:  math.MaxInt64,
		Floating: math.MaxFloat64,
		Pairs: []*message.TextNumberPair{
			{
				Text:   "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
				Number: math.MaxInt64,
			},
		},
	}

	small_jsonMessage = messageJson{
		Integer:  9,
		Floating: 1.1,
	}

	medium_jsonMessage = messageJson{
		Integer:  math.MaxInt32,
		Floating: math.MaxFloat32,
		Pairs: []*textNumberPair{
			{
				Text:   "Lorem ipsum",
				Number: math.MaxInt32,
			},
		},
	}

	big_jsonMessage = messageJson{
		Integer:  math.MaxInt64,
		Floating: math.MaxFloat64,
		Pairs: []*textNumberPair{
			{
				Text:   "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
				Number: math.MaxInt64,
			},
		},
	}

	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			if "ws://"+r.Host == r.Header.Get("Origin") {
				return true
			}
			return false
		},
	}
)

type messageJson struct {
	Integer  int               `json:"integer,omitempty"`
	Floating float64           `json:"float,omitempty"`
	Pairs    []*textNumberPair `json:"pairs,omitempty"`
}

type textNumberPair struct {
	Text   string `json:"text,omitempty"`
	Number int    `json:"number,omitempty"`
}

func wsHandlerProto(w http.ResponseWriter, r *http.Request) {

	wsProtoConnection, _ := upgrader.Upgrade(w, r, nil) // error ignored

	_ = wsProtoConnection.WriteMessage(websocket.BinaryMessage, data)

	_, _, _ = wsProtoConnection.ReadMessage()

	startProto(wsProtoConnection)
}

func wsHandlerJson(w http.ResponseWriter, r *http.Request) {
	wsJsonConnection, _ := upgrader.Upgrade(w, r, nil) // error ignored

	_ = wsJsonConnection.WriteJSON(big_jsonMessage)

	_ = wsJsonConnection.ReadJSON(&big_jsonMessage)

	startJson(wsJsonConnection)
}

func startProto(wsProtoConnection *websocket.Conn) {

	count = 0

	_ = wsProtoConnection.WriteMessage(websocket.BinaryMessage, data)

	var (
		receivedMessageBytes        []byte
		receivedMessageUnmarshalled = &message.Message{}
		messageToBeSent             []byte
	)

	start := time.Now()
	for {

		_, receivedMessageBytes, _ = wsProtoConnection.ReadMessage()

		_ = proto.Unmarshal(receivedMessageBytes, receivedMessageUnmarshalled)

		messageToBeSent, _ = proto.Marshal(receivedMessageUnmarshalled)

		_ = wsProtoConnection.WriteMessage(websocket.BinaryMessage, messageToBeSent)

		count++

		if count == times {
			end := time.Now()
			elapsedTime := end.Sub(start)

			_ = wsProtoConnection.Close()
			fmt.Printf("Protobufs elapsed time: %s\n", elapsedTime.String())
			return
		}
	}
}

func startJson(wsJsonConnection *websocket.Conn) {

	count = 0

	_ = wsJsonConnection.WriteJSON(small_jsonMessage)

	start := time.Now()
	for {

		_ = wsJsonConnection.ReadJSON(&small_jsonMessage)

		count++

		_ = wsJsonConnection.WriteJSON(small_jsonMessage)

		if count == times {
			_ = wsJsonConnection.Close()
			end := time.Now()
			elapsedTime := end.Sub(start)
			fmt.Printf("JSON elapsed time: %s\n", elapsedTime.String())
			count = 0
			return
		}
	}
}

func main() {
	//defer profile.Start().Stop()
	if len(os.Args) > 1 {
		if givenTimes, err := strconv.Atoi(os.Args[1]); err == nil {
			times = givenTimes
		}
	}

	data, _ = proto.Marshal(small_protoMessage)

	http.HandleFunc("/proto", wsHandlerProto)

	http.HandleFunc("/json", wsHandlerJson)

	_ = http.ListenAndServe("localhost:8080", nil)

}
