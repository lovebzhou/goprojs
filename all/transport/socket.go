package transport

import (
	"bufio"
	"io"
	"net"
)

type socketTransport struct {
	conn net.Conn
	rw   io.ReadWriter
	br   *bufio.Reader
	bw   *bufio.Writer
}

func NewSocketTransport() Transport {


