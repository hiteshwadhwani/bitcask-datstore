package pkg

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func RunInteractiveMode() error {
	fmt.Println("Welcome to interactive mode. Press Ctrl+C to exit.")

	// Set up signal handling for Ctrl+C
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Start a goroutine to handle the interrupt signal
	go func() {
		<-sigChan
		fmt.Println("\nExiting interactive mode...")
		os.Exit(0)
	}()

	// Start the interactive loop
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if scanner.Scan() {
			input := scanner.Text()
			// Process the input here
			fmt.Printf("You entered: %s\n", input)
		}
	}
}
