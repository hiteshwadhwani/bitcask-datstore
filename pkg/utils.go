package pkg

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

type InteractiveSession struct {
	commands []Command
}

type Command struct {
	Name string
	Run  func(args []string)
}

func NewInteractiveSession() *InteractiveSession {
	interactiveSession := &InteractiveSession{
		commands: make([]Command, 0),
	}

	return interactiveSession
}

func (i *InteractiveSession) StartInteractiveSession() {
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

	defer close(sigChan)

	// Start the interactive loop
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if scanner.Scan() {
			input := scanner.Text()
			i.HandleCommands(input)
		}
	}
}

func (i *InteractiveSession) HandleCommands(text string) {
	// use cobra to handle commands
	command, userArgs := strings.Split(text, " ")[0], strings.Split(text, " ")[1:]

	flag := false

	if command == "" {
		fmt.Println("Invalid command")
		return
	}

	for _, cmd := range i.commands {
		if cmd.Name == command {
			flag = true
			cmd.Run(userArgs)
		}
	}

	if !flag {
		fmt.Println("Invalid command")
	}
}

func (i *InteractiveSession) AddCommand(command Command) {
	i.commands = append(i.commands, command)
}
