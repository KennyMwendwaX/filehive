package p2p

import "net"

// Peer is an interface that represents the remote node.
type Peer interface {
	net.Conn
	Send([]byte) error
	CloseStream()
}

// Transport is anything that handles the communication
// between nodes in the network.
type Transport interface {
	ListenAndAccept() error
}
