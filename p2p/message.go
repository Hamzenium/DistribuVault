package p2p

// message represents any arbitrary data that is being sent over the each transport between nodes in the network.
type Message struct {
	payload []byte
}
