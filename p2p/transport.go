package p2p

// peer is interface that represents the remote node.
type Peer interface {
	Close() error
}

// transport is anything that handles the communcation
// between the nodes in the network
// can be of form (TCP, UDP, websockets)
type Transport interface {
	ListenAndAccept() error
	Consume() <-chan RPC
}
