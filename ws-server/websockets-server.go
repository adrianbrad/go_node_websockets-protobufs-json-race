package main

import (
	"fmt"
	"github.com/pkg/profile"
	"net/http"
	"time"

	"websockets-protobufs-json-race/ws-server/message"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
)

var times = 1000000

var count = 0

var data []byte

var jsonMessage messageJson

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		if "http://"+r.Host == r.Header.Get("Origin") {
			return true
		}
		return false
	},
}

type messageJson struct {
	Integer  int               `json:"integer"`
	Floating float64           `json:"float"`
	Pairs    []*textNumberPair `json:"pairs"`
}

type textNumberPair struct {
	Text   string `json:"text"`
	Number int    `json:"number"`
}

func wsHandlerProto(w http.ResponseWriter, r *http.Request) {

	conn, _ := upgrader.Upgrade(w, r, nil) // error ignored

	_ = conn.WriteMessage(websocket.BinaryMessage, data)

	_, msg, _ := conn.ReadMessage()

	_ = conn.WriteMessage(websocket.BinaryMessage, data)

	start := time.Now()
	for {
		//startReadingBytesTime := time.Now()

		_, msg, _ = conn.ReadMessage()

		//endReadingBytesTime := time.Now()

		//elapsedTimeReadingBytes := endReadingBytesTime.Sub(startReadingBytesTime)
		//fmt.Print("Protobufs reading elapsed time: " + string(elapsedTimeReadingBytes.String()+"\n"))

		receivedMessage := &message.Message{}
		//
		_ = proto.Unmarshal(msg, receivedMessage)

		count++

		toSend, _ := proto.Marshal(receivedMessage)

		_ = conn.WriteMessage(websocket.BinaryMessage, toSend)

		count++

		if count == times {
			_ = conn.Close()
			end := time.Now()
			elapsedTime := end.Sub(start)
			fmt.Printf("Protobufs elapsed time: %s\n", elapsedTime.String())
			count = 0
			return
		}
	}
}

func wsHandlerJson(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil) // error ignored

	_ = conn.WriteJSON(jsonMessage)

	msg := messageJson{}

	_ = conn.ReadJSON(&msg)

	_ = conn.WriteJSON(jsonMessage)
	start := time.Now()
	for {
		//startReadingJsonTime := time.Now()

		_ = conn.ReadJSON(&msg)

		//endReadingJsonTime := time.Now()
		//elapsedTimeReadingJson := endReadingJsonTime.Sub(startReadingJsonTime)

		//fmt.Print("Json reading elapsed time: " + string(elapsedTimeReadingJson.String()) + "\n")

		count++

		//if err := conn.WriteJSON(msg); err != nil {
		//	return
		//}

		_ = conn.WriteJSON(msg)

		if count == times {
			_ = conn.Close()
			end := time.Now()
			elapsedTime := end.Sub(start)
			fmt.Printf("JSON elapsed time: %s\n", elapsedTime.String())
			count = 0
			return
		}
	}

}

func functionElapsedTime(function string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", function, time.Since(start))
	}
}

func main() {
	defer profile.Start().Stop()

	data, _ = composeBinaryPayload()

	jsonMessage = composeJsonPayload()

	http.HandleFunc("/proto", wsHandlerProto)

	http.HandleFunc("/json", wsHandlerJson)

	_ = http.ListenAndServe("localhost:8080", nil)

}

func composeBinaryPayload() ([]byte, error) {
	return proto.Marshal(&message.Message{
		Integer: 12,
		//	Floating: 12123412.1251313561,
		//	Pairs:  []*message.TextNumberPair {
		//		{
		//			Text: "First Pair AFsamfnm,fqnwrmkllrqrqw;",
		//			Number: 123456789124512,
		//		},			{
		//			Text: "First Pair AFsamfnm,fqnwrmkllrqrqw;",
		//			Number: 123456789124512,
		//		},			{
		//			Text: "First Pair AFsamfnm,fqnwrmkllrqrqw;",
		//			Number: 123456789124512,
		//		},			{
		//			Text: "First Pair AFsamfnm,fqnwrmkllrqrqw;",
		//			Number: 123456789124512,
		//		},			{
		//			Text: "First Pair AFsamfnm,fqnwrmkllrqrqw;",
		//			Number: 123456789124512,
		//		},			{
		//			Text: "First Pair AFsamfnm,fqnwrmkllrqrqw;",
		//			Number: 123456789124512,
		//		},
		//	},
	})
}

func composeJsonPayload() messageJson {
	return messageJson{
		Integer: 12,
		//floating:   12124515.121516125,
		//pairs:  []*textNumberPair{
		//	{
		//		text:   "First Pair afjqwjkrwnjkfwnq",
		//		number: 123456789124512,
		//	},			{
		//		text:   "First Pair afjqwjkrwnjkfwnq",
		//		number: 123456789124512,
		//	},			{
		//		text:   "First Pair afjqwjkrwnjkfwnq",
		//		number: 123456789124512,
		//	},			{
		//		text:   "First Pair afjqwjkrwnjkfwnq",
		//		number: 123456789124512,
		//	},			{
		//		text:   "First Pair afjqwjkrwnjkfwnq",
		//		number: 123456789124512,
		//	},
		//},
	}
}
