package wsclient

import (
	"net"
)

type Conn struct {
	*wsclient.websocket
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
