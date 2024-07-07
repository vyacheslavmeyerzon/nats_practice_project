package goroutines

import (
	"fmt"  // Import the package for formatted input/output
	"sync" // Import the package for synchronizing goroutines
)

// Function to launch 10 goroutines, each printing numbers from 1 to 10
func LaunchGoroutines() {
	// Print a message about launching 10 goroutines
	fmt.Println("\n--- Launching 10 goroutines ---")

	// Create a WaitGroup variable to wait for all goroutines to complete
	var wg sync.WaitGroup

	// Launch 10 goroutines
	for i := 1; i <= 10; i++ {
		// Increment the WaitGroup counter by 1 before launching each goroutine
		wg.Add(1)

		// Launch a goroutine
		go func(id int) {
			// Decrement the WaitGroup counter by 1 after the goroutine completes
			defer wg.Done()

			// Print numbers from 1 to 10
			for j := 1; j <= 10; j++ {
				fmt.Printf("Goroutine %d: %d\n", id, j)
			}
		}(i)
	}

	// Wait for all goroutines to complete
	wg.Wait()
}

// Function to send data to a channel and print it from another goroutine
func ChannelExample() {
	// Print a message about launching the channel example
	fmt.Println("\n--- Sending data to a channel and receiving data from a channel ---")

	// Create a channel for integers
	ch := make(chan int)

	// Create a WaitGroup variable to wait for all goroutines to complete
	var wg sync.WaitGroup

	// Launch a goroutine to send data to the channel
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			ch <- i // Send data to the channel
		}
		close(ch) // Close the channel after sending all data
	}()

	// Launch a goroutine to receive data from the channel and print it
	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range ch {
			fmt.Println("Received:", val) // Receive data from the channel until it is closed
		}
	}()

	// Wait for all goroutines to complete
	wg.Wait()
}

// Function to demonstrate buffered channels
func BufferedChannelExample() {
	// Print a message about launching the buffered channels example
	fmt.Println("\n--- Example of using buffered channels ---")

	// Create a buffered channel with a capacity of 5
	ch := make(chan int, 5)

	// Create a WaitGroup variable to wait for all goroutines to complete
	var wg sync.WaitGroup

	// Launch a goroutine to send data to the buffered channel
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			fmt.Println("Sending:", i)
			ch <- i // Send data to the buffered channel
			fmt.Println("Sent:", i)
		}
		close(ch) // Close the channel after sending all data
	}()

	// Launch a goroutine to receive data from the buffered channel and print it
	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range ch {
			fmt.Println("Received from buffered channel:", val)
		}
	}()

	// Wait for all goroutines to complete
	wg.Wait()
}

// Function to demonstrate the select statement
func SelectExample() {
	// Print a message about launching the select statement example
	fmt.Println("\n--- Example of using the select statement ---")

	// Create two channels
	ch1 := make(chan int)
	ch2 := make(chan int)

	// Create a WaitGroup variable to wait for all goroutines to complete
	var wg sync.WaitGroup

	// Launch a goroutine to send data to the first channel
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 5; i++ {
			ch1 <- i // Send data to the first channel
		}
		close(ch1) // Close the first channel after sending all data
	}()

	// Launch a goroutine to send data to the second channel
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 6; i <= 10; i++ {
			ch2 <- i // Send data to the second channel
		}
		close(ch2) // Close the second channel after sending all data
	}()

	// Launch a goroutine to receive data from both channels using select
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case val, ok := <-ch1:
				if ok {
					fmt.Println("Received from ch1:", val) // Process data from the first channel
				} else {
					ch1 = nil // Set the channel to nil to avoid further reads
				}
			case val, ok := <-ch2:
				if ok {
					fmt.Println("Received from ch2:", val) // Process data from the second channel
				} else {
					ch2 = nil // Set the channel to nil to avoid further reads
				}
			}

			// Break the loop if both channels are closed
			if ch1 == nil && ch2 == nil {
				break
			}
		}
	}()

	// Wait for all goroutines to complete
	wg.Wait()
}
