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

type TCPTransportOpts struct {
	ListenAddress string
	Handshake     HandshakeFunc
	Decoder       Decoder
}

type TCPTransport struct {
	TCPTransportOpts
	listener net.Listener

	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	listener, err := net.Listen("tcp", t.ListenAddress)
	if err != nil {
		return err
	}

	t.listener = listener

	go t.acceptConnections()

	return nil
}

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

	if err := t.Handshake(peer); err != nil {
		conn.Close()
		fmt.Println("Error performing handshake:", err)
		return
	}

	msg := &Message{}
	for {
		if err := t.Decoder.Decode(conn, msg); err != nil {
			fmt.Println("Error decoding message:", err)
			continue
		}

		msg.From = conn.RemoteAddr()

		fmt.Println("Received message from:", msg.From, "with payload:", string(msg.Payload))
	}

}
