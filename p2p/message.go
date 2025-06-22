package p2p

import "net"

// RPC holds any data being sent over each transport network
type RPC struct {
	From    net.Addr
	Payload []byte
}
