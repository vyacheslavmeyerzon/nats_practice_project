package main

import (
	// Import the package for working with goroutines
	"nats_practice/goroutines"
	// Import the package for working with NATS
	"nats_practice/nats_basic"
)

func main() {
	// Launch the function for creating and executing goroutines
	goroutines.LaunchGoroutines()

	// Launch the function for working with channels
	goroutines.ChannelExample()

	// Launch the function for working with buffered channels
	goroutines.BufferedChannelExample()

	// Launch the function for working with the select operator
	goroutines.SelectExample()

	// Launch the Request-Reply NATS example
	nats_basic.RequestReplyExample()

	// Launch the Pub-Sub NATS example
	nats_basic.PubSubExample()

	// Launch the Queue Subscribe NATS example
	nats_basic.QueueSubscribeExample()

	// Launch the JetStream example
	nats_basic.JetStreamExample()
}
