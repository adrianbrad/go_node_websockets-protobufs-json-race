package benchmarks

import (
	"fmt"
	"math"
	"net/http"

	"github.com/gorilla/websocket"
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

type wsServerJson struct {
	jsonMessage *messageJson

	connection *websocket.Conn

	init chan bool
}

func NewJson() websocketsServer {
	ws := wsServerJson{}

	ws.init = make(chan bool)

	ws.initConnection()
	return &ws
}

func (wss *wsServerJson) initConnection() {
	http.HandleFunc("/json", wss.wsHandlerJson)
	go http.ListenAndServe("localhost:8080", nil)
}

func (wss *wsServerJson) wsHandlerJson(w http.ResponseWriter, r *http.Request) {
	wss.connection, _ = upgrader.Upgrade(w, r, nil)
	fmt.Println("Client connected")

	_ = wss.connection.WriteJSON(wss.jsonMessage)

	fmt.Println("Send verification message")

	_ = wss.connection.ReadJSON(wss.jsonMessage)
	fmt.Println("Received verification message")

	wss.init <- true
}

func (wss wsServerJson) WaitForInitializaion() {
	<-wss.init
}

func (wss *wsServerJson) SendAndReceiveMessage() {
	_ = wss.connection.WriteJSON(wss.jsonMessage)

	_ = wss.connection.ReadJSON(wss.jsonMessage)
}

func (wss *wsServerJson) SetMessageSize(size messageSize) {
	switch size {
	case small:
		wss.jsonMessage = &messageJson{
			Integer:  9,
			Floating: 1.1,
		}
	case medium:
		wss.jsonMessage = &messageJson{
			Integer:  math.MaxInt32,
			Floating: math.MaxFloat32,
			Pairs: []*textNumberPair{
				{
					Text:   "Lorem ipsum",
					Number: math.MaxInt32,
				},
			},
		}
	case big:
		wss.jsonMessage = &messageJson{
			Integer:  math.MaxInt64,
			Floating: math.MaxFloat64,
			Pairs: []*textNumberPair{
				{
					Text:   "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
					Number: math.MaxInt64,
				},
			},
		}
	}
}

func (wss *wsServerJson) CloseConnection() {
	fmt.Println("Closed conn")
	wss.connection.Close()
}
