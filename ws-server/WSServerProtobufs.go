package main

import (
	"fmt"
	"net/http"
	message "websockets-protobufs-json-race/proto"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
)

type wsServerProtobufs struct {
	smallMessage *message.Message

	mediumMessage message.Message

	bigMessage message.Message

	protoMessageBytes []byte

	receivedMessageBytes []byte

	receivedMessage *message.Message

	connection *websocket.Conn
}

func (wss *wsServerProtobufs) InitConnection() {
	wss.smallMessage = &message.Message{
		Integer:  9,
		Floating: 1.1,
	}
	wss.protoMessageBytes, _ = proto.Marshal(wss.smallMessage)

	http.HandleFunc("/proto", wss.wsHandlerProto)
	_ = http.ListenAndServe("localhost:8080", nil)
}

func (wss *wsServerProtobufs) wsHandlerProto(w http.ResponseWriter, r *http.Request) {
	wss.connection, _ = upgrader.Upgrade(w, r, nil)
	fmt.Println("got conn")

	_ = wss.connection.WriteMessage(websocket.BinaryMessage, wss.protoMessageBytes)
	fmt.Println("sent mess")

	_, _, _ = wss.connection.ReadMessage() //just used for initialization and caching
	fmt.Println("received mess")
}

func (wss *wsServerProtobufs) SendAndReceiveMessage() {
	wss.protoMessageBytes, _ = proto.Marshal(&message.Message{
		Integer:  9,
		Floating: 1.1,
	})

	fmt.Println(wss.connection)
	// fmt.Println(wss.protoMessageBytes)
	// _ = wss.connection.WriteMessage(websocket.BinaryMessage, wss.protoMessageBytes)

	// _, wss.receivedMessageBytes, _ = wss.connection.ReadMessage()
	// _ = proto.Unmarshal(wss.receivedMessageBytes, wss.receivedMessage)

	fmt.Println("received go")

	// wss.connection.Close()
}
