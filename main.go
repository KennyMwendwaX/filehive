package main

import (
	"log"

	"github.com/KennyMwendwaX/filehive/p2p"
)

func main() {
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
	}
	transport := p2p.NewTCPTransport(tcpOpts)

	if err := transport.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}
