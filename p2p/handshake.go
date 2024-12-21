package p2p

type HandshakeFunc func(Peer) error

func NopHandshakeFunc(Peer) error {
	return nil
}
