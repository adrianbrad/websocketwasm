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
	t := testevents.Start(t_, "TestWSHandshakeSucess", true)
	defer t.Done()

	_, err := websocketwasm.Dial(getWSBaseURL() + "echo")
	if err != nil {
		t.Fatalf("Unexpected error during handshake: %s", err.Error())
	}
}

func TestWSHandshakeFailed(t_ *testing.T) {
	t := testevents.Start(t_, "TestWSHandshakeFailed", true)
	defer t.Done()

	wsConn, err := websocketwasm.Dial(getWSBaseURL() + "invalid-enpoint")

	if err == nil {
		wsConn.Close()
		t.Fatalf("Got no error, but expected an error in opening the WebSocket.")
	}

	t.Logf("WebSocket failed to open: %s", err)
}

func TestWSSendAndReceiveTextMessageSucess(t_ *testing.T) {
	t := testevents.Start(t_, "TestSendTextMessageSuccess", true)
	defer t.Done()
	messageToBeSent := `{"wtf":1}`
	wsConn, _ := websocketwasm.Dial(getWSBaseURL() + "echo")

	wsConn.WriteString(messageToBeSent)
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
	wsConn, _ := websocketwasm.Dial(getWSBaseURL() + "echo")

	fmt.Println(wsConn.Write(messageToBeSent))

	mes := make([]byte, 10)
	n, err := wsConn.Read(mes)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(mes[:n], messageToBeSent) {
		t.Fatalf("Received message: %s not equal to expected message: %s", string(mes[:n]), messageToBeSent)
	}
}

func TestMain(t_ *testing.T) {
	t := testevents.Start(t_, "Add", true)
	defer t.Done()

	connect := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		go func() {
			ws, _ := websocketwasm.Dial(getWSBaseURL() + "echo")
			wsJSV := js.ValueOf(ws)
			wsJSV.Set("sendMessage", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
				go func() {
					ws.WriteString(args[0].String())
				}()
				return nil
			}))
			js.Global().Set("usr", wsJSV)
		}()
		return nil
	})

	js.Global().Set("add", connect)

	select {}
}
