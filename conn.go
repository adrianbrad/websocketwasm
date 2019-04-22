package websocketwasm

import (
	"net"

	"github.com/adrianbrad/websocketwasm/wsclient"
)

type Conn struct {
	*wsclient.Websocket
}

func Dial(url string) (net.Conn, error) {
	return nil, nil
}

func (c *Conn) Read() {
}

func (c *Conn) Write() {

}

func (c *Conn) Close() {

}
