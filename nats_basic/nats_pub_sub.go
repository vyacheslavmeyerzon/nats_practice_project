package nats_basic

import (
	"fmt"  // Import the package for formatted input/output
	"log"  // Import the package for logging errors
	"time" // Import the package for working with time

	"github.com/nats-io/nats.go" // Import the package for working with NATS
)

// Function to setup a NATS publisher and subscriber
func PubSubExample() {
	// Print a message about launching the Pub-Sub example
	fmt.Println("\n--- Example of using Pub-Sub NATS ---")

	// Connect to NATS server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err) // Log an error if the connection fails
	}
	defer nc.Close() // Close the connection when the function completes

	fmt.Println("Connected to NATS server") // Message about successful connection

	// Setup a subscriber to receive messages
	_, err = nc.Subscribe("updates", func(m *nats.Msg) {
		fmt.Printf("Received message: %s\n", string(m.Data)) // Print the received message
	})
	if err != nil {
		log.Fatal(err) // Log an error if setting up the subscriber fails
	}

	fmt.Println("Subscriber set up") // Message about successful subscriber setup

	// Allow some time for the subscriber to set up
	time.Sleep(1 * time.Second)

	// Publish a message
	err = nc.Publish("updates", []byte("Hello, World!"))
	if err != nil {
		log.Fatal(err) // Log an error if publishing the message fails
	}

	fmt.Println("Message published: Hello, World!") // Message about successful publication

	// Allow some time for the subscriber to receive the message
	time.Sleep(1 * time.Second)
}
