package wsclient

import "syscall/js"

type Websocket struct {
	js.Value
}

func NewWebsocket(url string) (ws *Websocket, err error) {
	defer handleErrorIfRaised(err)
	ws = &Websocket{js.Global().Get("WebSocket").New(url)}
	return
}

func (w *Websocket) onOpen(callback js.Func) {
	w.Call("addEventListener", "open", callback)
}

func (w *Websocket) onClose(callback js.Func) {
	w.Call("addEventListener", "close", callback)
}

func (w *Websocket) send(data js.Value) {
	w.Value.Call("send", data)
}

func (w *Websocket) onMessage(callback js.Func) {
	w.Call("addEventListener", "message", callback)
}

func (w *Websocket) close(callback js.Func) {
	w.Value.Call("close", 1000)
}

func handleErrorIfRaised(err error) {
	e := recover()
	if e == nil {
		return
	}
	if jsErr, ok := e.(*js.Error); ok && jsErr != nil {
		err = jsErr
	} else {
		panic(e)
	}
	return
}
