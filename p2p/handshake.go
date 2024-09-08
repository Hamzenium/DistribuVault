package p2p

type handshakeFunc func(any) error

func NOPHandShakeFunc(Peer) error { return nil }
