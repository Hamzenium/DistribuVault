package p2p

import (
	"fmt"
	"net"
	"sync"
)

// tcp peer represents the remote node over a tcp established connection
type TCPPeer struct {
	conn     net.Conn // the underlying connetion of the peer
	outbound bool     // if we dial and retirve a conn -> outbound == true, else if we accept conn-> outbound == false
}

// func indicates the start of a function definition.
// NewTCPPeer is the name of the function. new struct instances are named with a New prefix.
// conn net.Conn is the first parameter, representing an established network connection (net.Conn is an interface representing a generic network connection).
// outbound bool is the second parameter, a boolean value indicating whether the connection is outbound or not.
// *TCPPeer is the return type, which is a pointer to a TCPPeer struct.
func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

type TCPTransport struct {
	listenAddress string
	listener      net.Listener
	handshakeFunc HandshakeFunc

	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

// func NOPHandShakeFunc(any) error { return nil }

// this is the constructor function for the tcp transport
// NewTCPTransport is the name of the function. By convention, constructor functions in Go often start with New.
// listenAddr string is the parameter, which specifies the address to listen on.
// *TCPTransport is the return type of the function, indicating it returns a pointer to a TCPTransport struct.
func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		handshakeFunc: NOPHandShakeFunc,
		listenAddress: listenAddr,
	}
}

// func indicates a function.
// (t *TCPTransport) specifies that this is a method with a receiver t, which is a pointer to a TCPTransport struct.
// This means ListenAndAccept is a method associated with TCPTransport.
// ListenAndAccept is the name of the method.
// error is the return type, indicating that this method returns an error if something goes wrong.
func (t *TCPTransport) ListenAndAccept() error {

	//var err error declares a variable err of type error.
	//t.listener, err = net.Listen("tcp", t.listenAddress) attempts to start listening on the TCP network
	//at the address specified by t.listenAddress.
	// The net.Listen function creates a TCP listener.
	// If successful, it assigns the listener to t.listener. \
	//If it fails, err will contain the error.
	var err error
	t.listener, err = net.Listen("tcp", t.listenAddress)
	if err != nil {
		return err
	}
	//go starts a new goroutine, which is a lightweight concurrent function.
	// It calls t.startAcceptLoop(), presumably a method that handles accepting incoming connections in a loop.
	//This allows the server to handle connections concurrently without blocking the rest of the program.
	go t.startAcceptLoop()

	return nil
}

// func indicates a function.
// (t *TCPTransport) specifies that this is a method with the receiver t, a pointer to a TCPTransport struct.
// startAcceptLoop is the name of the method, and it doesn’t take any parameters or return anything.
func (t *TCPTransport) startAcceptLoop() {
	//This for loop without a condition runs indefinitely, creating an infinite loop.
	//This is common in server applications where continuous operation is required,
	// such as listening for incoming connections.
	//t.listener.Accept() waits for and accepts an incoming TCP connection.
	//If a connection is successfully accepted, it's assigned to the variable conn.
	//If an error occurs during acceptance, it's assigned to the variable err.
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept error: %s\n", err)
		}
		go t.handleConn(conn)
	}
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)
	if err := t.handshakeFunc(conn); err != nil {

	}
	lenDecodeError := 
	msg := &Temp{}
	for {
		if err := t.decoder.Decode(conn, msg) ; err != nil  {
			fmt.Printf("TCP error: %s\n", err)
			continue
		}
	}
	fmt.Printf(("new incoming connection %+v\n"), peer)
}
