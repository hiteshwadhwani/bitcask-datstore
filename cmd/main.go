// This is the implementation of the bitcask database papaer
// https://riak.com/assets/bitcask-intro.pdf

package main

import (
	"fmt"
	"hiteshwadhwani/bitcask-datstore.git/internal/bitcask"
	"hiteshwadhwani/bitcask-datstore.git/pkg"
	"os"
)

func main() {
	db, err := bitcask.NewDiskStore("bitcask.db")

	if err != nil {
		fmt.Printf("something went wrong: %v", err)
	}

	defer db.Close()
	defer os.Remove("bitcask.db")

	var commands = []pkg.Command{
		{
			Name: "set",
			Run: func(args []string) {
				if len(args) < 2 {
					fmt.Println("set command requires a key and a value")
					return
				}

				key, value := args[0], args[1]

				err := db.Set(key, value)
				if err != nil {
					fmt.Printf("error setting value: %v \n", err)
				}

				fmt.Printf("value set successfully: %v \n", value)
			},
		},
		{
			Name: "get",
			Run: func(args []string) {
				if len(args) < 1 {
					fmt.Println("get command requires a key")
					return
				}

				key := args[0]

				value, err := db.Get(key)

				if err != nil {
					fmt.Printf("error getting value: %v \n", err)
					return
				}

				fmt.Printf("value: %v \n", value)
			},
		},
		{
			Name: "delete",
			Run: func(args []string) {
				if len(args) < 1 {
					fmt.Println("delete command requires a key")
					return
				}

				key := args[0]

				err := db.Delete(key)
				if err != nil {
					fmt.Printf("error deleting value: %v \n", err)
					return
				}

				fmt.Println("value deleted successfully")
			},
		},
		{
			Name: "exit",
			Run: func(args []string) {
				os.Exit(0)
			},
		},
	}

	cliSession := pkg.NewInteractiveSession()

	for _, command := range commands {
		cliSession.AddCommand(command)
	}
	cliSession.StartInteractiveSession()

}
