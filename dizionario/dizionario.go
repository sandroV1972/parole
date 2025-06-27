package main

import "sync"

type dizionario struct {
	data map[string]string
	mu   sync.RWMutex
}

var instance *dizionario
var once sync.Once

func newDizionario() dizionario {
	once.Do(func() {
		instance = &dizionario{
			data: make(map[string]string),
		}
	})
	return *instance
}

func (d *dizionario) Reset() {
	d.mu.Lock()
	d.data = make(map[string]string)
	d.mu.Unlock()
}

// Set imposta un valore per una chiave
func (d *dizionario) Set(key, value string) {
	d.mu.Lock()
	d.data[key] = value
	d.mu.Unlock()
}

// Get ottiene un valore per una chiave
func (d *dizionario) Get(key string) string {
	d.mu.RLock()
	value := d.data[key]
	d.mu.RUnlock()
	return value
}
