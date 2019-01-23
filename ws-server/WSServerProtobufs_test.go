package main

import (
	"os/exec"
	"testing"
	"time"
)

func TestSendAndReceiveProto(b *testing.T) {
	ws := wsServerProtobufs{}
	ws.InitConnection()
	time.Sleep(1000 * time.Millisecond)
	cmd := exec.Command("node", "../ws-client/clientProto.js")
	// output, _ := cmd.CombinedOutput()
	cmd.Run()
	// _, _ = cmd.CombinedOutput()
	// fmt.Println(string(output))
	// for n := 0; n < 1; n++ {
	// ws.SendAndReceiveMessage()
	// }
}
