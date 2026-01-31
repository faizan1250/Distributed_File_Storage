package p2p

// peer is an interface that represents the remote node
type Peer interface{}

// Transport is anythig that handles the connection
// between the nodes in the network. This can be of
// the form (TCP, UDP, webSockets, ...)
type Transport interface {
	ListenAndAccept() error
}
