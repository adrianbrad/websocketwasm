package websocketwasm

import (
	"fmt"
	"syscall/js"

	"github.com/adrianbrad/websocketwasm/browser"
)

type WebSocket struct {
	*browser.WebSocket
	Send     chan []byte
	Received chan []byte

	onOpenFunc    js.Func
	onCloseFunc   js.Func
	onMessageFunc js.Func
}

func New(url string) (ws *WebSocket, err error) {
	wsb, err := browser.NewWebsocket(url)
	if err != nil {
		return
	}
	ws = &WebSocket{
		WebSocket: wsb,
		Send:      make(chan []byte),
		Received:  make(chan []byte),
	}
	ws.onOpenFunc = js.FuncOf(ws.onOpenListener)
	ws.onCloseFunc = js.FuncOf(ws.onCloseListener)
	ws.onMessageFunc = js.FuncOf(ws.onMessageListener)

	wsb.OnOpen(ws.onOpenFunc)
	wsb.OnClose(ws.onCloseFunc)

	return nil, nil
}

func (w *WebSocket) onOpenListener(this js.Value, args []js.Value) interface{} {
	fmt.Println("Open")
	return nil
}

func (w *WebSocket) onMessageListener(this js.Value, args []js.Value) interface{} {
	fmt.Println("Message")
	fmt.Println(args[0])
	return nil
}

func (w *WebSocket) onCloseListener(this js.Value, args []js.Value) interface{} {
	fmt.Println("Close")

	return nil
}
