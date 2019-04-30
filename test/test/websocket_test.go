package websocket_test

import (
	"bytes"
	"fmt"
	"syscall/js"
	"testing"

	"github.com/LinearZoetrope/testevents"
	"github.com/adrianbrad/websocketwasm"
)

func TestWSHandshakeSucess(t_ *testing.T) {
	t := testevents.Start(t_, "TestConnSuccess", true)
	defer t.Done()

	testFinished := make(chan struct{})
	ws, err := websocketwasm.New("ws://localhost:3000/ws/echo")
	ws.OnOpen(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if err != nil {
			t.Fatal(err.Error())
		}
		testFinished <- struct{}{}
		return nil
	}))
	<-testFinished
}

func TestWSSendAndReceiveTextMessageSucess(t_ *testing.T) {
	t := testevents.Start(t_, "TestSendTextMessageSuccess", true)
	defer t.Done()
	messageToBeSent := `{"wtf":1}`
	wsConn, _ := websocketwasm.New("ws://localhost:3000/ws/echo")
	wsConn.OnOpen(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		fmt.Println(wsConn.WriteString(messageToBeSent))
		return nil

	}))
	mes := make([]byte, 10)
	n, err := wsConn.Read(mes)
	if err != nil {
		t.Fatal(err)
	}

	if string(mes[:n]) != messageToBeSent {
		t.Fatalf("Received message: %s not equal to expected message: %s", string(mes[:n]), messageToBeSent)
	}
}

func TestWSSendAndReceiveBinaryMessageSucess(t_ *testing.T) {
	t := testevents.Start(t_, "TestSendBinaryMessageSuccess", true)
	defer t.Done()
	messageToBeSent := []byte(`{"wtf":1}`)
	wsConn, _ := websocketwasm.New("ws://localhost:3000/ws/echo")
	wsConn.OnOpen(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		fmt.Println(wsConn.Write(messageToBeSent))
		return nil
	}))

	mes := make([]byte, 10)
	n, err := wsConn.Read(mes)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(mes[:n], messageToBeSent) {
		t.Fatalf("Received message: %s not equal to expected message: %s", string(mes[:n]), messageToBeSent)
	}
}
