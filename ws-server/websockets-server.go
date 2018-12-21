package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"websockets-protobufs-json-race/proto"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
)

var (
	times        = 20000
	count        = 0
	data         []byte
	protoMessage = &message.Message{
		Integer:  12,
		Floating: 12.12,
	}

	jsonMessage = messageJson{
		Integer:  12,
		Floating: 12.12,
	}

	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			if "ws://"+r.Host == r.Header.Get("Origin") {
				return true
			}
			return false
		},
	}

	wsProtoConnection *websocket.Conn
	wsJsonConnection  *websocket.Conn
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

	wsProtoConnection, _ = upgrader.Upgrade(w, r, nil) // error ignored

	_ = wsProtoConnection.WriteMessage(websocket.BinaryMessage, data)

	_, _, _ = wsProtoConnection.ReadMessage()

	startProto()
}

func wsHandlerJson(w http.ResponseWriter, r *http.Request) {
	wsJsonConnection, _ = upgrader.Upgrade(w, r, nil) // error ignored

	_ = wsJsonConnection.WriteJSON(jsonMessage)

	_ = wsJsonConnection.ReadJSON(&jsonMessage)

	startJson()
}

func startProto() {

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

func startJson() {

	count = 0

	_ = wsJsonConnection.WriteJSON(jsonMessage)

	start := time.Now()
	for {

		_ = wsJsonConnection.ReadJSON(&jsonMessage)

		count++

		_ = wsJsonConnection.WriteJSON(jsonMessage)

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

	data, _ = proto.Marshal(protoMessage)

	http.HandleFunc("/proto", wsHandlerProto)

	http.HandleFunc("/json", wsHandlerJson)

	_ = http.ListenAndServe("localhost:8080", nil)

}
