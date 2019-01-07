package main

import (
	"fmt"
	"github.com/pkg/profile"
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

	times             = 20000
	count             = 0
	protoMessageBytes []byte
	jsonMessage       messageJson
	profileStart      interface{ Stop() }
	//profileStart profile.Profile
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

	_ = wsProtoConnection.WriteMessage(websocket.BinaryMessage, protoMessageBytes)

	_, _, _ = wsProtoConnection.ReadMessage()

	startProto(wsProtoConnection)
}

func wsHandlerJson(w http.ResponseWriter, r *http.Request) {
	wsJsonConnection, _ := upgrader.Upgrade(w, r, nil) // error ignored

	_ = wsJsonConnection.WriteJSON(jsonMessage)

	_ = wsJsonConnection.ReadJSON(&jsonMessage)

	startJson(wsJsonConnection)
}

func startProto(wsProtoConnection *websocket.Conn) {

	count = 0

	_ = wsProtoConnection.WriteMessage(websocket.BinaryMessage, protoMessageBytes)

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

	_ = wsJsonConnection.WriteJSON(jsonMessage)

	start := time.Now()
	for {

		_ = wsJsonConnection.ReadJSON(&jsonMessage)

		_ = wsJsonConnection.WriteJSON(jsonMessage)

		count++

		if count == times {
			_ = wsJsonConnection.Close()
			end := time.Now()
			elapsedTime := end.Sub(start)
			fmt.Printf("JSON elapsed time: %s\n\n", elapsedTime.String())
			count = 0
			profileStart.Stop()

			os.Exit(0)
			return
		}
	}
}

func main() {
	if len(os.Args) > 1 {
		if givenTimes, err := strconv.Atoi(os.Args[1]); err == nil {
			times = givenTimes
		}

		switch os.Args[2] {
		case "s":
			protoMessageBytes, _ = proto.Marshal(small_protoMessage)
			jsonMessage = small_jsonMessage
			fmt.Println("Small messages")
		case "m":
			protoMessageBytes, _ = proto.Marshal(medium_protoMessage)
			jsonMessage = medium_jsonMessage
			fmt.Println("Medium messages")
		case "b":
			protoMessageBytes, _ = proto.Marshal(big_protoMesage)
			jsonMessage = big_jsonMessage
			fmt.Println("Big messages")
		}
	} else {
		times = 20000
		jsonMessage = small_jsonMessage
		protoMessageBytes, _ = proto.Marshal(small_protoMessage)
	}

	fmt.Printf("%d times\n", times)

	http.HandleFunc("/proto", wsHandlerProto)

	http.HandleFunc("/json", wsHandlerJson)

	profileStart = profile.Start(profile.ProfilePath(os.Getenv("GOPATH") + "/pprofs"))

	_ = http.ListenAndServe("localhost:8080", nil)

}
