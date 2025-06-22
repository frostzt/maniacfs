package p2p

// Peer represents a remote node
type Peer interface {
	Close() error
}

// Transport handles comms between nodes
type Transport interface {
	ListenAndAccept() error
	Consume() <-chan RPC
}
