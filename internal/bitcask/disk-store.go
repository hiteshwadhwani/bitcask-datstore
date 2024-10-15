// DiskStore is a disk based store for the Bitcask datastore

package bitcask

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

const (
	defaultWhence = io.SeekStart
)

type keyDirEntry struct {
	timestamp uint32
	position  uint32
	size      uint32
}

type DiskStore struct {
	// file the datafile on disk
	// At a time only one file will be active
	file *os.File

	// my is a mutex to prevent race condition
	my *sync.Mutex

	// keyDir is a map of key-value pair
	// It will be used to get the position of the key-value pair in the file
	keyDir map[string]keyDirEntry

	// position of cursor in the file
	// will be use to write new key-value pair in the file
	writePosition int
}

func checkIfFileExists(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func NewDiskStore(filePath string) (*DiskStore, error) {
	d := &DiskStore{
		my:            &sync.Mutex{},
		keyDir:        make(map[string]keyDirEntry),
		writePosition: 0,
	}
	exists, err := checkIfFileExists(filePath)
	if err != nil {
		return nil, fmt.Errorf("error checking if file exists: %v", err)
	}

	if exists {
		err := d.initKeyDir(filePath)
		if err != nil {
			return nil, fmt.Errorf("error initializing key directory: %v", err)
		}
	}

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, fmt.Errorf("error opening data file: %v", err)
	}
	d.file = file

	return d, nil
}

// It will be called at startup to initialize the keyDir
func (d *DiskStore) initKeyDir(filePath string) error {
	file, err := os.Open(filePath)

	if err != nil {
		return fmt.Errorf("error opening data file: %v", err)
	}

	defer file.Close()

	for {
		header := make([]byte, headerSize)

		_, err := io.ReadFull(file, header)

		if err == io.EOF {
			break
		}

		if err != nil {
			return fmt.Errorf("error reading header: %v", err)
		}

		timestamp, keySize, valueSize := decodeHeader(header)

		key := make([]byte, keySize)
		value := make([]byte, valueSize)

		_, err = io.ReadFull(file, key)

		if err != nil {
			return fmt.Errorf("error reading key: %v", err)
		}

		_, err = io.ReadFull(file, value)

		if err != nil {
			return fmt.Errorf("error reading value: %v", err)
		}

		totalSize := headerSize + keySize + valueSize

		d.keyDir[string(key)] = keyDirEntry{
			timestamp: timestamp,
			position:  uint32(d.writePosition) + totalSize,
			size:      totalSize,
		}
	}

	return nil

}

func (d *DiskStore) Get(key string) (string, error) {
	entry, ok := d.keyDir[key]

	if !ok {
		return "", fmt.Errorf("key not found")
	}

	entryBuffer := make([]byte, entry.size)

	_, err := d.file.Seek(int64(entry.position), defaultWhence)
	if err != nil {
		return "", fmt.Errorf("error seeking to position: %v", err)
	}

	_, err = io.ReadFull(d.file, entryBuffer)
	if err == io.EOF {
		return "", fmt.Errorf("error reading value: %v", err)
	}

	_, value := decodeKeyValue(entryBuffer)

	return value, nil
}

func (d *DiskStore) Put(key string, value string) error {
	timestamp := uint32(time.Now().Unix())

	totalSize, kv := encodeKeyValue(timestamp, key, value)

	_, err := d.file.Seek(int64(d.writePosition), defaultWhence)
	if err != nil {
		return fmt.Errorf("error seeking to position: %v", err)
	}

	_, err = d.file.Write(kv)

	if err != nil {
		return fmt.Errorf("error writing key-value pair: %v", err)
	}

	d.keyDir[key] = keyDirEntry{
		timestamp: timestamp,
		position:  uint32(d.writePosition),
		size:      uint32(totalSize),
	}

	d.writePosition += totalSize

	return nil
}

func (d *DiskStore) Close() {
	d.file.Close()
}
