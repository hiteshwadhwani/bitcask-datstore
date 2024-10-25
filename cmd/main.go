// This is the implementation of the bitcask database papaer
// https://riak.com/assets/bitcask-intro.pdf

package main

import "hiteshwadhwani/bitcask-datstore.git/internal/bitcask"

func main() {
	// db, err := bitcask.NewDiskStore("bitcask.db")

	// if err != nil {
	// 	fmt.Printf("something went wrong: %v", err)
	// }

	// defer db.Close()
	// defer os.Remove("bitcask.db")

	// if err != nil {
	// 	fmt.Printf("error creating disk store: %v", err)
	// }

	// err = db.Set("name", "hitesh")

	// if err != nil {
	// 	fmt.Printf("error putting value: %v", err)
	// }

	// value, err := db.Get("name")
	// if err != nil {
	// 	fmt.Printf("error getting value: %v", err)
	// }

	// fmt.Println(value)

	// err = db.Delete("name")
	// if err != nil {
	// 	fmt.Printf("error deleting value: %v", err)
	// }

	// value, err = db.Get("name")
	// if err != nil {
	// 	fmt.Printf("error getting value: %v", err)
	// }

	// fmt.Println(value)

	bitcask.Execute()
}
