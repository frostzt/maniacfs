package p2p

import "net"

// Message holds any data being sent over each transport network
type Message struct {
	From    net.Addr
	Payload []byte
}
