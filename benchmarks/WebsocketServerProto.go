package benchmarks

import (
	"fmt"
	"math"
	"net/http"
	message "websockets-protobufs-json-race/proto"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
)

type wsServerProtobufs struct {
	message *message.Message

	protoMessageBytes []byte

	receivedMessageBytes []byte

	receivedMessage *message.Message

	connection *websocket.Conn

	init chan bool
}

func NewProto() websocketsServer {
	ws := wsServerProtobufs{}

	ws.init = make(chan bool)
	ws.receivedMessage = &message.Message{}

	ws.initConnection()
	return &ws
}

func (wss *wsServerProtobufs) initConnection() {
	http.HandleFunc("/proto", wss.wsHandlerProto)
	go http.ListenAndServe("localhost:8080", nil)
}

func (wss *wsServerProtobufs) wsHandlerProto(w http.ResponseWriter, r *http.Request) {
	wss.connection, _ = upgrader.Upgrade(w, r, nil)
	fmt.Println("Client connected")

	_ = wss.connection.WriteMessage(websocket.BinaryMessage, wss.protoMessageBytes)
	fmt.Println("Send verification message")

	_, _, _ = wss.connection.ReadMessage() //just used for initialization and caching
	fmt.Println("Received verification message")

	wss.init <- true
}

func (wss wsServerProtobufs) WaitForInitializaion() {
	<-wss.init
}

func (wss *wsServerProtobufs) SendAndReceiveMessage() {
	wss.protoMessageBytes, _ = proto.Marshal(wss.message)
	_ = wss.connection.WriteMessage(websocket.BinaryMessage, wss.protoMessageBytes)
	_, wss.receivedMessageBytes, _ = wss.connection.ReadMessage()
	_ = proto.Unmarshal(wss.receivedMessageBytes, wss.receivedMessage)
}

func (wss *wsServerProtobufs) SetMessageSize(size messageSize) {
	switch size {
	case small:
		wss.message = &message.Message{
			Integer:  9,
			Floating: 1.1,
		}
	case medium:
		wss.message = &message.Message{
			Integer:  math.MaxInt32,
			Floating: math.MaxFloat32,
			Pairs: []*message.TextNumberPair{
				{
					Text:   "Lorem ipsum",
					Number: math.MaxInt32,
				},
			},
		}
	case big:
		wss.message = &message.Message{
			Integer:  math.MaxInt64,
			Floating: math.MaxFloat64,
			Pairs: []*message.TextNumberPair{
				{
					Text:   "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
					Number: math.MaxInt64,
				},
			},
		}
	}
	wss.protoMessageBytes, _ = proto.Marshal(wss.message)
}

func (wss *wsServerProtobufs) CloseConnection() {
	wss.connection.Close()
}
