/*
 * Copyright (c) 2024 Luca Fröschke
 */

package datastore

import (
	"bytes"
	"encoding/gob"
	"sync"
)

type DataStore struct {
	sync.RWMutex
	store   map[string]any
	plugins map[string][]any
}

var instance *DataStore
var once sync.Once

func GetInstance() *DataStore {
	once.Do(func() {
		instance = &DataStore{
			store:   make(map[string]any),
			plugins: make(map[string][]any),
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

func (ds *DataStore) Delete(key string) {
	ds.Lock()
	defer ds.Unlock()
	delete(ds.store, key)
}

func (ds *DataStore) Clear() {
	ds.Lock()
	defer ds.Unlock()
	ds.store = make(map[string]any)
}

func (ds *DataStore) GetAllData() map[string]any {
	ds.RLock()
	defer ds.RUnlock()

	dataCopy := make(map[string]any)
	for key, value := range ds.store {
		dataCopy[key] = value
	}
	return dataCopy
}

func (ds *DataStore) GetSize() float64 {
	ds.RLock()
	defer ds.RUnlock()

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	err := enc.Encode(ds.store)
	if err != nil {
		return 0
	}

	size := float64(buf.Len()) / 1024.0
	return size
}

func (ds *DataStore) PushPluginData(pluginName string, data any) {
	ds.Lock()
	defer ds.Unlock()
	ds.plugins[pluginName] = append(ds.plugins[pluginName], data)
}

func (ds *DataStore) GetPluginData(pluginName string) ([]any, bool) {
	ds.RLock()
	defer ds.RUnlock()
	data, ok := ds.plugins[pluginName]
	return data, ok
}

func (ds *DataStore) ClearPluginData(pluginName string) {
	ds.Lock()
	defer ds.Unlock()
	ds.plugins[pluginName] = make([]any, 0)
}

func (ds *DataStore) GetAllPluginData() map[string][]any {
	ds.RLock()
	defer ds.RUnlock()

	dataCopy := make(map[string][]any)
	for key, value := range ds.plugins {
		dataCopy[key] = value
	}
	return dataCopy
}

func (ds *DataStore) ClearAllPluginData() {
	ds.Lock()
	defer ds.Unlock()
	ds.plugins = make(map[string][]any)
}
