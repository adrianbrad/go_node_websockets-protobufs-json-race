package main

type WebsocketsServer interface {
	InitConnection()
	SendAndReceiveMessage()
}
