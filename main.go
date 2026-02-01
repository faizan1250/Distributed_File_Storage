package main

import (
	"log"

	"github.com/faizan1250/Distributed_File_Storage/p2p"
)

func main() {
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		Decoder:       p2p.NOPDecoder{},
		HandshakeFunc: p2p.NOPHandshakefunc,
	}
	tr := p2p.NewTCPTransport(tcpOpts)
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select {}
}
