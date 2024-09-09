package p2p

import "net"

// message represents any arbitrary data that is being sent over the each transport between nodes in the network.
type RPC struct {
	From    net.Addr
	Payload []byte
}
