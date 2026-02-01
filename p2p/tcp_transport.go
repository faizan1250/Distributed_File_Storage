package p2p

import (
	"fmt"
	"log/slog"
	"net"
	"sync"
)

// TCPPeer represents the remote node over a tcp established connection
type TCPPeer struct {
	// conn is the underlying connection of the peer
	conn net.Conn

	// if we dial and retrieve a conn => outbound ==true
	// if we accept and retrieve a conn => outbound ==false
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

type TCPTransportOpts struct {
	ListenAddr    string
	HandshakeFunc HandshakeFunc
	Decoder       Decoder
}
type TCPTransport struct {
	Listener net.Listener
	TCPTransportOpts
	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
		peers:            make(map[net.Addr]Peer),
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	ln, err := net.Listen("tcp", t.ListenAddr)
	if err != nil {
		return err
	}
	t.Listener = ln
	go t.startAcceptLoop()
	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.Listener.Accept()
		if err != nil {
			slog.Error("transport accepts error", "err", err)
			continue
		}
		fmt.Printf("new incoming peer : %+v\n", conn)
		go t.handleConn(conn)

	}
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)

	if err := t.HandshakeFunc(peer); err != nil {
		conn.Close()
		fmt.Printf("TCP Handshake error %v\n", err)
		return
	}

	msg := Message{}
	for {
		if err := t.Decoder.Decode(conn, &msg); err != nil {
			fmt.Printf("tcp msg read err : %v\n", err)
		}
		msg.From = conn.RemoteAddr()
		fmt.Printf("message=> from : %+v , message: %+v\n", msg.From, string(msg.Payload))
	}
}
