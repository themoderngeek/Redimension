package service

import "sync"

type Redimension struct {
	store map[string]string
	mu    sync.RWMutex
}

func NewRedimension() *Redimension {
	return &Redimension{
		store: make(map[string]string),
	}
}

func (r *Redimension) Set(key, value string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.store[key] = value
}

func (r *Redimension) Get(key string) (string, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	value, exists := r.store[key]
	return value, exists
}
