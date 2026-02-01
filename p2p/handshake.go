package p2p

import "errors"

var ErrInvalidHandshake = errors.New("invalid handshake")

type HandshakeFunc func(Peer) error

func NOPHandshakefunc(Peer) error {
	return nil
}
