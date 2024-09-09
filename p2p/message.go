package p2p

import "net"

// message represents any arbitrary data that is being sent over the each transport between nodes in the network.
type Message struct {
	From    net.Addr
	Payload []byte
}
