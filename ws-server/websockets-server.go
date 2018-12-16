package main

import (
	"fmt"
	"net/http"
	"time"

	"./message"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
)


var times = int32(10000)

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
	Times int32 `json:"times"`
	Pair  []*StringInt32NumberPairJson `json:"pairs"`
}

type StringInt32NumberPairJson struct {
	Str string `json:"string"`
	Int32Number int32 `json:"number"`
}

func wsHandlerProto(w http.ResponseWriter, r *http.Request) {

	conn, _ := upgrader.Upgrade(w, r, nil) // error ignored

	start := time.Now()

	_ = conn.WriteMessage(websocket.BinaryMessage, data)

	for {
			_, msg, _ := conn.ReadMessage()

			receivedMessage := &message.Message{}

			_ = proto.Unmarshal(msg, receivedMessage)

			if receivedMessage.Times < times {
				count ++;

				incrementMessage(receivedMessage)

				toSend, _ := proto.Marshal(receivedMessage)

				if err := conn.WriteMessage(websocket.BinaryMessage, toSend); err != nil {
					return
				}

			} else {
				_ = conn.Close()
				count = 0
				t := time.Now()
				elapsed := t.Sub(start)
				fmt.Print("Protobufs elapsed time: " + string(elapsed.String() + "\n"))
				return
			}
	}
}

func wsHandlerJson(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil) // error ignored

	start := time.Now()

	_ = conn.WriteJSON(jsonMessage)

	for {
		msg := messageJson{}
		_ = conn.ReadJSON(&msg)

		if msg.Times < times {
			count ++;

			msg.Times ++;

			if err := conn.WriteJSON(msg); err != nil {
				return
			}

		} else {
			_ = conn.Close()
			count = 0
			t := time.Now()
			elapsed := t.Sub(start)
			fmt.Print("Json elapsed time: " + string(elapsed.String() + "\n"))
			return
		}
	}

}

func incrementMessage(message *message.Message) {
	message.Times ++
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
			Str: "First Pair",
			Int32Number: int32(12345),
		}

	pairs = append(pairs, &pair)

	msg := &message.Message{
		Times: 0,
		Pair:  pairs,
	}

	return proto.Marshal(msg);
}

func composeJsonPayload() messageJson{
	pairs := make([]*StringInt32NumberPairJson, 0)

	pair := StringInt32NumberPairJson{
		Str: "First Pair",
		Int32Number: int32(12345),
	}

	pairs = append(pairs, &pair)

	msg := messageJson{
		Times: 0,
		Pair: pairs,
	}

	return msg
}
