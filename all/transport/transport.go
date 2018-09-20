package transport

import (
	"crypto/tls"
	"io"
)

// Type represents a transport type
type Type int

const (
	// Socket represents a TCP transport type.
	Socket Type = iota + 1

	// WebSocket reprsents a websocket type.
	WebSocket
)

func (t Type) String() string {
	switch t {
	case Socket:
		return "socket"
	case WebSocket:
		return "websocket"
	}
	return ""
}

type Transport interface {
	io.ReadWriteCloser

	Type() Type

	StartTLS(cfg *tls.Config, asClient bool)
	
}
