package bitcask

import (
	"encoding/binary"
)

// headerSize is the size of the header in the data file
const headerSize = 12

func encodeHeader(timeStamp, keySize, valueSize uint32) []byte {
	header := make([]byte, headerSize)

	binary.LittleEndian.PutUint32(header[0:4], timeStamp)
	binary.LittleEndian.PutUint32(header[4:8], keySize)
	binary.LittleEndian.PutUint32(header[8:12], valueSize)

	return header
}

func decodeHeader(header []byte) (uint32, uint32, uint32) {
	timeStamp := binary.LittleEndian.Uint32(header[0:4])
	keySize := binary.LittleEndian.Uint32(header[4:8])
	valueSize := binary.LittleEndian.Uint32(header[8:12])

	return timeStamp, keySize, valueSize
}

func encodeKeyValue(timestamp uint32, key string, value string) (int, []byte) {
	// len(key) not returns the len of string but it returns the len of bytes in go
	header := encodeHeader(timestamp, uint32(len(key)), uint32(len(value)))

	kv := make([]byte, 0, headerSize+len([]byte(key))+len([]byte(value)))

	kv = append(kv, header...)
	kv = append(kv, []byte(key)...)
	kv = append(kv, []byte(value)...)

	return len(kv), kv
}

func decodeKeyValue(kv []byte) (string, string) {
	_, keySize, valueSize := decodeHeader(kv[:headerSize])

	key := string(kv[headerSize : headerSize+keySize])
	value := string(kv[headerSize+keySize : headerSize+keySize+valueSize])

	return key, value
}
