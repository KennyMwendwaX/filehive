package p2p

// Peer is a node in the network.
type Peer interface {
}

// Transport is anything that handles the communication
// between nodes in the network.
type Transport interface {
	ListenAndAccept() error
}
