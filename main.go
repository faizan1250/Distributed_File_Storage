package main

import (
	"fmt"
	"log"

	"github.com/faizan1250/Distributed_File_Storage/p2p"
)

func main() {
	onPeer := func(p p2p.Peer) error {
		p.Close()
		return nil
	}
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		Decoder:       p2p.NOPDecoder{},
		HandshakeFunc: p2p.NOPHandshakefunc,
		OnPeer:        onPeer,
	}
	tr := p2p.NewTCPTransport(tcpOpts)
	go func() {
		for {
			msg := <-tr.Consume()
			fmt.Printf("msg: %+v\n", msg)
		}
	}()
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select {}
}
