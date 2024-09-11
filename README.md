TCP-based Peer-to-Peer (P2P) Networking System
Overview
This document provides a detailed explanation of a TCP-based peer-to-peer (P2P) networking system implemented in Go. The system includes components for managing TCP connections, performing handshakes, decoding messages, and handling communication between peers.

Components
1. TCPTransport
Purpose: Manages TCP connections, listens for incoming connections, performs handshakes, and uses decoders to interpret messages.

Fields:

ListenAddr: The address where the transport listens for incoming connections (e.g., ":3000").
HandshakeFunc: A function used to perform a handshake with peers.
Decoder: An implementation of the Decoder interface for decoding incoming messages.
Methods:

NewTCPTransport(opts TCPTransportOpts) *TCPTransport: Constructor that initializes a new TCPTransport with the given options.
ListenAndAccept() error: Starts the transport system, listening for incoming connections and handling them.
startAcceptLoop(): Continuously accepts incoming connections and delegates them to handleConn.
handleConn(conn net.Conn): Manages an individual TCP connection, performs handshakes, and processes messages.
2. TCPPeer
Purpose: Represents a remote node in the network, associated with a TCP connection.

Fields:

conn: The underlying TCP connection.
outbound: Boolean indicating if the connection is outbound (true if initiated by this peer, false if accepted).
Constructor:

NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer: Creates a new TCPPeer instance with the specified connection and direction.
3. HandshakeFunc
Purpose: A function type used to perform handshakes between peers.

Usage: Can be customized to implement specific handshake logic.

Example Implementation:

NOPHandShakeFunc(any) error { return nil }: A no-operation handshake function that does nothing.
4. Decoder
Purpose: Interface for decoding data from an io.Reader into an RPC object.

Methods:

Decode(io.Reader, *RPC) error: Decodes data from the reader into the RPC struct.
5. GOBDecoder
Purpose: Implements the Decoder interface using Go's gob package for binary decoding.

Methods:

Decode(r io.Reader, msg *RPC) error: Uses gob.NewDecoder to decode data from the reader.
6. DefaultDecoder
Purpose: Implements the Decoder interface with basic read operations.

Methods:

Decode(r io.Reader, msg *RPC) error: Reads data from the reader into a buffer and assigns it to the RPC payload.
7. RPC
Purpose: Represents a message exchanged between peers.

Fields:

From: The address of the peer sending the message.
Payload: The actual data being transmitted.
8. Interfaces
Peer Interface:

Represents a connected node in the network.
Currently empty but can be extended with peer-specific methods.
Transport Interface:

Represents any communication mechanism (e.g., TCP, UDP, Websockets) that can listen for and accept connections.
Methods:
ListenAndAccept() error
