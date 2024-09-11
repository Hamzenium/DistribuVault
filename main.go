package main

import (
	"fmt"
	"log"

	"github.com/Hamzenium/Distributed-CAS-Storage/p2p"
)

// main initializes and runs the TCP-based peer-to-peer transport system.
func main() {
	// Configure the TCP transport options
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",              // Address to listen for incoming connections (port 3000)
		HandshakeFunc: p2p.NOPHandShakeFunc, // No-operation handshake function (placeholder)
		Decoder:       p2p.DefaultDecoder{}, // Default decoder for reading messages
	}

	// Create a new TCP transport instance with the specified options
	tr := p2p.NewTCPTransport(tcpOpts)

	// Start a goroutine to handle and print incoming messages
	go func() {
		for {
			// Receive a message from the transport's consume channel
			msg := <-tr.Consume() // Assuming Consume() returns a channel for incoming messages
			// Print the received message
			fmt.Printf("%+v\n", msg) // Fixed the newline character from /n to \n
		}
	}()

	// Start listening for incoming connections
	if err := tr.ListenAndAccept(); err != nil {
		// Log and terminate if there's an error while starting the transport
		log.Fatal(err)
	}

	// Block the main goroutine indefinitely to keep the program running
	select {}
}
