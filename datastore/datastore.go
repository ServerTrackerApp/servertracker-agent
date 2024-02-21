/*
 * Copyright (c) 2024 Luca Fröschke
 */

package datastore

import (
	"encoding/json"
	"sync"
)

type DataStore struct {
	sync.RWMutex
	store map[string]json.RawMessage
}

var instance *DataStore
var once sync.Once

func GetInstance() *DataStore {
	once.Do(func() {
		instance = &DataStore{
			store: make(map[string]json.RawMessage),
		}
	})
	return instance
}

func (ds *DataStore) Set(key string, value json.RawMessage) {
	ds.Lock()
	defer ds.Unlock()
	ds.store[key] = value
}

func (ds *DataStore) Get(key string) (json.RawMessage, bool) {
	ds.RLock()
	defer ds.RUnlock()
	value, ok := ds.store[key]
	return value, ok
}
