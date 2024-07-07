package nats_basic

import (
	"fmt"  // Import the package for formatted input/output
	"log"  // Import the package for logging errors
	"sync" // Import the package for synchronizing goroutines
	"time" // Import the package for working with time

	"github.com/nats-io/nats.go" // Import the package for working with NATS
)

// Function to setup NATS queue subscribers
func QueueSubscribeExample() {
	// Print a message about launching the Queue Subscribe example
	fmt.Println("\n--- Example of using Queue Subscribe NATS ---")

	// Connect to NATS server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err) // Log an error if the connection fails
	}
	defer nc.Close() // Close the connection when the function completes

	fmt.Println("Connected to NATS server") // Message about successful connection

	// Create a WaitGroup variable to wait for all subscribers to complete
	var wg sync.WaitGroup

	// Setup the first queue subscriber
	wg.Add(1)
	go func() {
		defer wg.Done()
		_, err := nc.QueueSubscribe("tasks", "worker", func(m *nats.Msg) {
			fmt.Printf("Worker 1 received: %s\n", string(m.Data)) // Print the received message
		})
		if err != nil {
			log.Fatal(err) // Log an error if setting up the subscriber fails
		}
		fmt.Println("Worker 1 subscribed to queue") // Message about successful subscription
	}()

	// Setup the second queue subscriber
	wg.Add(1)
	go func() {
		defer wg.Done()
		_, err := nc.QueueSubscribe("tasks", "worker", func(m *nats.Msg) {
			fmt.Printf("Worker 2 received: %s\n", string(m.Data)) // Print the received message
		})
		if err != nil {
			log.Fatal(err) // Log an error if setting up the subscriber fails
		}
		fmt.Println("Worker 2 subscribed to queue") // Message about successful subscription
	}()

	// Allow some time for the subscribers to set up
	time.Sleep(1 * time.Second)

	// Publish messages
	for i := 1; i <= 5; i++ {
		err := nc.Publish("tasks", []byte(fmt.Sprintf("Task %d", i)))
		if err != nil {
			log.Fatal(err) // Log an error if publishing the message fails
		}
		fmt.Printf("Published message: Task %d\n", i) // Message about successful publication
	}

	// Allow some time for the subscribers to receive the messages
	time.Sleep(1 * time.Second)

	// Wait for all subscribers to complete
	wg.Wait()

	fmt.Println("All messages received and processed") // Message about successful processing of all messages
}
