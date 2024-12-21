package main

import (
	"log"

	"github.com/KennyMwendwaX/file-hive/p2p"
)

func main() {
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddress: ":3000",
		Handshake:     p2p.NOPHandshakeFunc,
		Decoder:       p2p.GOBDecoder{},
	}
	transport := p2p.NewTCPTransport(tcpOpts)

	if err := transport.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}
