package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./test/")))

	http.HandleFunc("/ws/echo", func(w http.ResponseWriter, r *http.Request) {
		u := websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
		wsConn, err := u.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println(err)
			return
		}

		for {
			opcode, m, err := wsConn.ReadMessage()
			if err != nil {
				return
			}
			wsConn.WriteMessage(opcode, m)
		}
	})

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}

func handshakeSucces(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	u := websocket.Upgrader{}
	return u.Upgrade(w, r, nil)
}
