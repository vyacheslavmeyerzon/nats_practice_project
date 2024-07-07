package nats_basic

import (
	"fmt"  // Import the package for formatted input/output
	"log"  // Import the package for logging errors
	"time" // Import the package for working with time

	"github.com/nats-io/nats.go" // Import the package for working with NATS
)

// Function to setup JetStream stream and publish messages
func JetStreamExample() {
	// Print a message about launching the JetStream example
	fmt.Println("\n--- Example of using JetStream NATS ---")

	// Connect to NATS server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err) // Log an error if the connection fails
	}
	defer nc.Close() // Close the connection when the function completes

	fmt.Println("Connected to NATS server") // Message about successful connection

	// Create JetStream context
	js, err := nc.JetStream()
	if err != nil {
		log.Fatal(err) // Log an error if creating the context fails
	}

	fmt.Println("JetStream context created") // Message about successful context creation

	// Check JetStream availability
	accountInfo, err := js.AccountInfo()
	if err != nil {
		log.Fatal("Error fetching JetStream account info: ", err) // Log an error if fetching account info fails
	}

	fmt.Printf("JetStream is available: %+v\n", accountInfo) // Print account info for JetStream

	// Define stream configuration
	streamConfig := &nats.StreamConfig{
		Name:      "ORDERS",             // Stream name
		Subjects:  []string{"orders.*"}, // Subjects for the stream
		Storage:   nats.FileStorage,     // File storage type
		Retention: nats.LimitsPolicy,    // Retention policy for messages
	}

	// Add stream
	streamInfo, err := js.AddStream(streamConfig)
	if err != nil {
		log.Fatal("Error adding stream: ", err) // Log an error if adding the stream fails
	}

	fmt.Println("Stream added: ", streamInfo.Config.Name) // Message about successful stream addition

	// Publish messages to the stream
	for i := 1; i <= 5; i++ {
		subject := fmt.Sprintf("orders.%d", i) // Define the subject of the message
		message := fmt.Sprintf("Order %d", i)  // Define the message
		ack, err := js.Publish(subject, []byte(message))
		if err != nil {
			log.Fatal("Error publishing message: ", err) // Log an error if publishing the message fails
		}
		fmt.Printf("Published message: %s to subject: %s, ack: %v\n", message, subject, ack) // Message about successful publication
	}

	// Allow some time for messages to be processed
	time.Sleep(2 * time.Second)

	// Create a pull-based consumer
	consumerConfig := &nats.ConsumerConfig{
		Durable:   "ORDER_CONSUMER",       // Durable consumer
		AckPolicy: nats.AckExplicitPolicy, // Explicit acknowledgment policy
	}

	consumerInfo, err := js.AddConsumer("ORDERS", consumerConfig)
	if err != nil {
		log.Fatal("Error adding consumer: ", err) // Log an error if adding the consumer fails
	}

	fmt.Println("Consumer added: ", consumerInfo.Name) // Message about successful consumer addition

	// Subscribe to the consumer
	sub, err := js.PullSubscribe("orders.*", "ORDER_CONSUMER")
	if err != nil {
		log.Fatal("Error subscribing to consumer: ", err) // Log an error if subscribing to the consumer fails
	}

	fmt.Println("Subscribed to the consumer") // Message about successful subscription

	// Fetch messages from the consumer with increased timeout
	msgs, err := sub.Fetch(5, nats.MaxWait(20*time.Second))
	if err != nil {
		log.Fatal("Error fetching messages: ", err) // Log an error if fetching messages fails
	}

	for _, msg := range msgs {
		fmt.Printf("Received message: %s\n", string(msg.Data)) // Print the received message
		msg.Ack()                                              // Acknowledge the message
	}

	fmt.Println("Messages fetched and acknowledged") // Message about successful message fetching and acknowledgment

	// Object store operations
	fmt.Println("\n--- Working with the object store ---")
	objectStoreExample(js)

	// Key-value store operations
	fmt.Println("\n--- Working with the key-value store ---")
	keyValueStoreExample(js)

	// Filtered subject consumer
	fmt.Println("\n--- Filtering subjects for consumers ---")
	filteredSubjectConsumer(js)

	// Consumer with ack wait and max deliver settings
	fmt.Println("\n--- Configuring consumers with ack wait and max deliver ---")
	consumerWithAckWaitAndMaxDeliver(js)
}

// Function to demonstrate object store operations
func objectStoreExample(js nats.JetStreamContext) {
	// Create an object store
	objStoreConfig := &nats.ObjectStoreConfig{
		Bucket: "MY_BUCKET", // Object store bucket name
	}

	objStore, err := js.CreateObjectStore(objStoreConfig)
	if err != nil {
		log.Fatal("Error creating object store: ", err) // Log an error if creating the object store fails
	}

	fmt.Println("Object store created") // Message about successful object store creation

	// Put an object in the store
	_, err = objStore.PutBytes("my_object", []byte("This is a test object"))
	if err != nil {
		log.Fatal("Error putting object in store: ", err) // Log an error if putting the object fails
	}

	fmt.Println("Object stored successfully") // Message about successful object storage

	// Get the object from the store
	obj, err := objStore.GetBytes("my_object")
	if err != nil {
		log.Fatal("Error getting object from store: ", err) // Log an error if getting the object fails
	}

	fmt.Printf("Retrieved object: %s\n", string(obj)) // Print the retrieved object

	// Delete the object from the store
	err = objStore.Delete("my_object")
	if err != nil {
		log.Fatal("Error deleting object from store: ", err) // Log an error if deleting the object fails
	}

	fmt.Println("Object deleted successfully") // Message about successful object deletion
}

// Function to demonstrate key-value store operations
func keyValueStoreExample(js nats.JetStreamContext) {
	// Create a key-value store
	kvStoreConfig := &nats.KeyValueConfig{
		Bucket: "MY_KV_BUCKET", // Key-value store bucket name
	}

	kvStore, err := js.CreateKeyValue(kvStoreConfig)
	if err != nil {
		log.Fatal("Error creating key-value store: ", err) // Log an error if creating the key-value store fails
	}

	fmt.Println("Key-Value store created") // Message about successful key-value store creation

	// Put a key-value pair in the store
	_, err = kvStore.Put("my_key", []byte("This is a test value"))
	if err != nil {
		log.Fatal("Error putting key-value pair in store: ", err) // Log an error if putting the key-value pair fails
	}

	fmt.Println("Key-Value pair stored successfully") // Message about successful key-value pair storage

	// Get the value from the store
	kvEntry, err := kvStore.Get("my_key")
	if err != nil {
		log.Fatal("Error getting key-value pair from store: ", err) // Log an error if getting the key-value pair fails
	}

	fmt.Printf("Retrieved key-value pair: key=%s, value=%s\n", kvEntry.Key(), string(kvEntry.Value())) // Print the retrieved key-value pair

	// Delete the key-value pair from the store
	err = kvStore.Delete("my_key")
	if err != nil {
		log.Fatal("Error deleting key-value pair from store: ", err) // Log an error if deleting the key-value pair fails
	}

	fmt.Println("Key-Value pair deleted successfully") // Message about successful key-value pair deletion
}

// Function to demonstrate filtered subject consumer
func filteredSubjectConsumer(js nats.JetStreamContext) {
	// Create a consumer with filtered subjects
	consumerConfig := &nats.ConsumerConfig{
		Durable:       "FILTERED_CONSUMER",    // Durable consumer
		FilterSubject: "orders.1",             // Subject filter
		AckPolicy:     nats.AckExplicitPolicy, // Explicit acknowledgment policy
	}

	consumerInfo, err := js.AddConsumer("ORDERS", consumerConfig)
	if err != nil {
		log.Fatal("Error adding filtered consumer: ", err) // Log an error if adding the filtered consumer fails
	}

	fmt.Println("Filtered consumer added: ", consumerInfo.Name) // Message about successful filtered consumer addition

	// Subscribe to the filtered consumer
	sub, err := js.PullSubscribe("orders.1", "FILTERED_CONSUMER")
	if err != nil {
		log.Fatal("Error subscribing to filtered consumer: ", err) // Log an error if subscribing to the filtered consumer fails
	}

	fmt.Println("Subscribed to the filtered consumer") // Message about successful subscription

	// Fetch messages from the consumer with increased timeout
	msgs, err := sub.Fetch(5, nats.MaxWait(20*time.Second))
	if err != nil {
		log.Fatal("Error fetching messages from filtered consumer: ", err) // Log an error if fetching messages from the filtered consumer fails
	}

	for _, msg := range msgs {
		fmt.Printf("Received message from filtered consumer: %s\n", string(msg.Data)) // Print the received message from the filtered consumer
		msg.Ack()                                                                     // Acknowledge the message
	}

	fmt.Println("Messages from filtered consumer fetched and acknowledged") // Message about successful fetching and acknowledgment of messages from the filtered consumer
}

// Function to demonstrate consumer with ack wait and max deliver settings
func consumerWithAckWaitAndMaxDeliver(js nats.JetStreamContext) {
	// Create a consumer with ack wait and max deliver settings
	consumerConfig := &nats.ConsumerConfig{
		Durable:    "ACK_WAIT_CONSUMER",    // Durable consumer
		AckPolicy:  nats.AckExplicitPolicy, // Explicit acknowledgment policy
		AckWait:    10 * time.Second,       // Ack wait time
		MaxDeliver: 5,                      // Max delivery attempts
	}

	consumerInfo, err := js.AddConsumer("ORDERS", consumerConfig)
	if err != nil {
		log.Fatal("Error adding consumer with ack wait and max deliver: ", err) // Log an error if adding the consumer fails
	}

	fmt.Println("Consumer with ack wait and max deliver added: ", consumerInfo.Name) // Message about successful consumer addition

	// Subscribe to the consumer with ack wait and max deliver settings
	sub, err := js.PullSubscribe("orders.*", "ACK_WAIT_CONSUMER")
	if err != nil {
		log.Fatal("Error subscribing to consumer with ack wait and max deliver: ", err) // Log an error if subscribing to the consumer fails
	}

	fmt.Println("Subscribed to the consumer with ack wait and max deliver") // Message about successful subscription

	// Fetch messages from the consumer with increased timeout
	msgs, err := sub.Fetch(5, nats.MaxWait(20*time.Second))
	if err != nil {
		log.Fatal("Error fetching messages from consumer with ack wait and max deliver: ", err) // Log an error if fetching messages fails
	}

	for _, msg := range msgs {
		fmt.Printf("Received message from consumer with ack wait and max deliver: %s\n", string(msg.Data)) // Print the received message from the consumer
		msg.Ack()                                                                                          // Acknowledge the message
	}

	fmt.Println("Messages from consumer with ack wait and max deliver fetched and acknowledged") // Message about successful fetching and acknowledgment of messages from the consumer
}
