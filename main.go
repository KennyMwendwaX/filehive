package main

import (
	"log"

	"github.com/KennyMwendwaX/file-hive/p2p"
)

func main() {
	transport := p2p.NewTCPTransport(":3000")

	if err := transport.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}
