package p2p

import (
	"fmt"
	"net"
	"sync"
)

// TCPPeer represents the remote node in a TCP connection.
type TCPPeer struct {
	conn     net.Conn
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

type TCPTransport struct {
	listenAddress string
	listener      net.Listener
	handshake     HandshakeFunc
	decoder       Decoder

	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(listenAddress string) *TCPTransport {
	return &TCPTransport{
		handshake:     NopHandshakeFunc,
		listenAddress: listenAddress,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	listener, err := net.Listen("tcp", t.listenAddress)
	if err != nil {
		return err
	}

	t.listener = listener

	go t.acceptConnections()

	return nil
}

type Temp struct{}

func (t *TCPTransport) acceptConnections() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
		}

		fmt.Println("Handling connection from:", conn.RemoteAddr())

		go t.handleConnection(conn)
	}
}

func (t *TCPTransport) handleConnection(conn net.Conn) {
	peer := NewTCPPeer(conn, true)

	if err := t.handshake(conn); err != nil {
		fmt.Println("Error performing handshake:", err)
		return
	}

	msg := &Temp{}
	for {
		if err := t.decoder.Decode(conn, msg); err != nil {
			fmt.Println("Error decoding message:", err)
			continue
		}

	}

}
