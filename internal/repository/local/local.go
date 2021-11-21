package local

import (
	"errors"
	"sync"
)

var (
	ErrKeyNotFound  = errors.New("key not found")
	ErrKeyNotUnique = errors.New("key not unique")
)

type LocalStorage struct {
	db   map[string]interface{}
	mu   *sync.RWMutex
	bufK []string
	bufV []interface{}
}

func NewLocalStorage() LocalStorage {
	return LocalStorage{
		db:   make(map[string]interface{}),
		mu:   &sync.RWMutex{},
		bufK: make([]string, 0, 100),
		bufV: make([]interface{}, 0, 100),
	}
}

func (s LocalStorage) Set(key string, value interface{}) error {
	if _, ok := s.Get(key); ok {
		return ErrKeyNotUnique
	}
	s.mu.Lock()
	s.db[key] = value
	s.mu.Unlock()
	return nil
}

func (s LocalStorage) Get(key string) (interface{}, bool) {
	s.mu.RLock()
	value, ok := s.db[key]
	s.mu.RUnlock()
	return value, ok
}

func (s LocalStorage) Delete(key string) error {
	if _, ok := s.Get(key); !ok {
		return ErrKeyNotFound
	}
	s.mu.Lock()
	delete(s.db, key)
	s.mu.Unlock()
	return nil
}

func (s LocalStorage) Keys() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	s.bufK = s.bufK[:0]
	for k := range s.db {
		s.bufK = append(s.bufK, k)
	}
	return s.bufK
}

func (s LocalStorage) Values() []interface{} {
	s.mu.RLock()
	defer s.mu.RUnlock()
	s.bufV = s.bufV[:0]
	for _, v := range s.db {
		s.bufV = append(s.bufV, v)
	}
	return s.bufV
}
