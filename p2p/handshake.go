package p2p

import "errors"

// ErrInvalidHandshake is an error that occurs when the handshake between
// the local and remote nodes cannot be established successfully.
var ErrInvalidHandshake = errors.New("invalid Handshake")

// HandshakeFunc defines a function type for handshakes that takes a Peer as
// input and returns an error. This is used to standardize handshake functions
// across the P2P package.
type HandshakeFunc func(Peer) error

// NOPHandShakeFunc is a no-operation (no-op) handshake function that takes a Peer
// and returns nil. It serves as a placeholder or default implementation when no
// actual handshake logic is needed.
func NOPHandShakeFunc(Peer) error {
	return nil
}
