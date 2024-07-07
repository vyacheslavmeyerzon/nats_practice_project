package nats_basic

import (
	"fmt"  // Import the package for formatted input/output
	"log"  // Import the package for logging errors
	"time" // Import the package for working with time

	"github.com/nats-io/nats.go" // Import the package for working with NATS
)

// Function to setup a NATS server connection and perform request-reply
func RequestReplyExample() {
	// Print a message about launching the Request-Reply example
	fmt.Println("\n--- Example of using Request-Reply NATS ---")

	// Connect to NATS server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err) // Log an error if the connection fails
	}
	defer nc.Close() // Close the connection when the function completes

	fmt.Println("Connected to NATS server") // Message about successful connection

	// Setup a subscriber to reply to requests
	_, err = nc.Subscribe("request", func(m *nats.Msg) {
		fmt.Printf("Received request: %s\n", string(m.Data)) // Print the received request
		m.Respond([]byte("response"))                        // Send a response to the request
	})
	if err != nil {
		log.Fatal(err) // Log an error if setting up the subscriber fails
	}

	fmt.Println("Subscriber set up to respond to requests") // Message about successful subscriber setup

	// Allow some time for the subscriber to set up
	time.Sleep(1 * time.Second)

	// Send a request and wait for a reply
	msg, err := nc.Request("request", []byte("hello"), 2*time.Second)
	if err != nil {
		log.Fatal(err) // Log an error if sending the request fails
	}

	fmt.Printf("Received reply: %s\n", string(msg.Data)) // Print the received reply
}
