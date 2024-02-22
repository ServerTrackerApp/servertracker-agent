/*
 * Copyright (c) 2024 Luca Fröschke
 */

package datastore

import (
	"sync"
)

type DataStore struct {
	sync.RWMutex
	store map[string]any
}

var instance *DataStore
var once sync.Once

func GetInstance() *DataStore {
	once.Do(func() {
		instance = &DataStore{
			store: make(map[string]any),
		}
	})
	return instance
}

func (ds *DataStore) Set(key string, value any) {
	ds.Lock()
	defer ds.Unlock()
	ds.store[key] = value
}

func (ds *DataStore) Get(key string) (any, bool) {
	ds.RLock()
	defer ds.RUnlock()
	value, ok := ds.store[key]
	return value, ok
}
