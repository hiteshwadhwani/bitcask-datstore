// This is the implementation of the bitcask database papaer
// https://riak.com/assets/bitcask-intro.pdf

package main

import (
	"fmt"
	"sync"
)

// BitcaskHandle is a handle to a Bitcask database
// Only one process can have read_write access to the datastore at a time
type BitcaskHandle struct {
	directoryName string

	// Read write lock to synchronize access to the datastore
	readWriteLock *sync.RWMutex
}

// Opts can be read_write or read_only
type BitcaskOption string

const (
	ReadOnly  BitcaskOption = "read_only"
	ReadWrite BitcaskOption = "read_write"
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

type Bitcask interface {
	// Will create or open existing datastore with the additional options

	// Opts can be read_write or read_only
	Open(directoryName string, Opts BitcaskOption) (*BitcaskHandle, error)

	// Put will put a key value pair into the database
	Put(handle *BitcaskHandle, key []byte, value []byte) error

	// Get will get a value for a key from the database
	Get(handle *BitcaskHandle, key []byte) ([]byte, error)

	// ListKeys will list all the keys in the database
	ListKeys(handle *BitcaskHandle) ([]string, error)

	// Delete will delete a key value pair from the database
	// Will add a tombstone to the key and value
	Delete(handle *BitcaskHandle, key []byte) error

	// Merge several data files into a single one within a bitcask datastore and produce a hintfile for faster startup
	Merge(handle *BitcaskHandle)

	// Close will close the database and flush all pending writes
	Close(handle *BitcaskHandle) error
}

func main() {
	fmt.Println("Hello, World!")
}
