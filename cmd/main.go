// This is the implementation of the bitcask database papaer
// https://riak.com/assets/bitcask-intro.pdf

package main

import (
	"fmt"
	"hiteshwadhwani/bitcask-datstore.git/internal/bitcask"
	"os"
)

// type BitcaskOptions struct{
// 	// Max number of data files to keep in the datastore
// 	maxDataFiles int

// 	// Max size of a data file
// 	maxDataFileSize int

// 	// Max size of a hint file
// 	maxHintFileSize int

// 	// Merge percentage
// 	mergePercentage int

// 	// Compression algorithm to use for the data file
// 	compressionAlgo string

// 	// Compression level to use for the data file
// 	compressionLevel int

// 	// Whether to use a bloom filter for the data file
// 	useBloomFilter bool

// 	// Bloom filter false positive rate
// 	bloomFilterFP float64

// 	// Whether to use a hint file for faster startup
// 	useHintFile bool

// 	// Whether to sync data to disk after each write
// 	syncWrites bool

// 	// Whether to use a merge worker to merge data files periodically
// 	useMergeWorker bool

// 	// Merge worker interval
// 	mergeWorkerInterval time.Duration
// }

func main() {
	db, err := bitcask.NewDiskStore("bitcask.db")

	defer db.Close()
	defer os.Remove("bitcask.db")

	if err != nil {
		fmt.Printf("error creating disk store: %v", err)
	}

	err = db.Put("name", "hitesh")

	if err != nil {
		fmt.Printf("error putting value: %v", err)
	}

	value, err := db.Get("name")
	if err != nil {
		fmt.Printf("error getting value: %v", err)
	}

	fmt.Println(value)
}
