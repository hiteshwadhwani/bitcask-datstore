package bitcask

import "sync"

type KeyLock map[string]*sync.RWMutex

func NewKeyLock() KeyLock {
	return make(KeyLock)
}

func (k KeyLock) GetLock(key string) *sync.RWMutex {
	lock, exists := k[key]
	if !exists {
		lock = &sync.RWMutex{}
		k[key] = lock
	}
	return lock
}
