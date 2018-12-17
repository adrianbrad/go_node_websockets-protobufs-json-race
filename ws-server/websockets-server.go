package main

import (
	"fmt"
	"net/http"
	"time"

	"./message"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
)

var times = int32(1)

var count = int32(0)

var data []byte

var jsonMessage messageJson

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		if r.Host == "localhost:8080" {
			return true
		}
		return false
	},
}

type messageJson struct {
	Times int32                        `json:"times"`
	Pair  []*StringInt32NumberPairJson `json:"pairs"`
}

type StringInt32NumberPairJson struct {
	Str         string `json:"string"`
	Int32Number int32  `json:"number"`
}

func wsHandlerProto(w http.ResponseWriter, r *http.Request) {

	conn, _ := upgrader.Upgrade(w, r, nil) // error ignored

	_ = conn.WriteMessage(websocket.BinaryMessage, data)

	//var msg = make([]byte, 10)
	for {
		startReadingBytesTime := time.Now()

		_, msg, _ := conn.ReadMessage() // FIRST TIME THIS IS SO SLOW

		endReadingBytesTime := time.Now()

		elapsedTimeReadingBytes := endReadingBytesTime.Sub(startReadingBytesTime)

		receivedMessage := &message.Message{}

		_ = proto.Unmarshal(msg, receivedMessage)

		fmt.Print("Protobufs reading elapsed time: " + string(elapsedTimeReadingBytes.String()+"\n"))

		count++

		toSend, _ := proto.Marshal(receivedMessage)

		if err := conn.WriteMessage(websocket.BinaryMessage, toSend); err != nil {
			return
		}

		if count == times {
			_ = conn.Close()
			//t := time.Now()
			count = 0
			//elapsed := t.Sub(start)
			//fmt.Print("Protobufs elapsed time: " + string(elapsed.String() + "\n"))
			return
		}
	}
}

func wsHandlerJson(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil) // error ignored

	_ = conn.WriteJSON(jsonMessage)

	for {
		msg := messageJson{}

		startReadingJsonTime := time.Now()

		_ = conn.ReadJSON(&msg)

		endReadingJsonTime := time.Now()
		elapsedTimeReadingJson := endReadingJsonTime.Sub(startReadingJsonTime)

		fmt.Print("Json reading elapsed time: " + string(elapsedTimeReadingJson.String()) + "\n")

		count++

		if err := conn.WriteJSON(msg); err != nil {
			return
		}

		if count == times {
			_ = conn.Close()
			//t := time.Now()
			//elapsed := t.Sub(start)
			count = 0
			//fmt.Print("Json elapsed time: " + string(elapsed.String() + "\n"))
			return
		}
	}

}

func main() {

	data, _ = composeBinaryPayload()

	jsonMessage = composeJsonPayload()

	http.HandleFunc("/proto", wsHandlerProto)

	http.HandleFunc("/json", wsHandlerJson)

	_ = http.ListenAndServe("localhost:8080", nil)
}

func composeBinaryPayload() ([]byte, error) {
	pairs := make([]*message.StringInt32NumberPair, 1)

	pair := message.StringInt32NumberPair{
		Str:         "First Pair",
		Int32Number: int32(1234567891),
	}

	pairs = append(pairs, &pair)

	msg := &message.Message{
		Times: 0,
		Pair:  pairs,
	}

	return proto.Marshal(msg)
}

func composeJsonPayload() messageJson {
	pairs := make([]*StringInt32NumberPairJson, 1)

	pair := StringInt32NumberPairJson{
		Str:         "First Pair",
		Int32Number: int32(1234567891),
	}

	pairs = append(pairs, &pair)

	msg := messageJson{
		Times: 0,
		Pair:  pairs,
	}

	return msg
}
