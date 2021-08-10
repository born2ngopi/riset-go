package memory

import (
	"errors"
	"sync"
)

type mem struct {
	mu      sync.Mutex
	storage map[string]interface{}
}

type Memory interface {
	Set(key string, value interface{}) error
	Get(key string) (interface{}, error)
	Delete(key string) error
}

func New() Memory {
	return &mem{
		storage: make(map[string]interface{}),
	}
}

func (m *mem) Set(key string, value interface{}) error {

	m.mu.Lock()
	{
		m.storage[key] = value
	}
	m.mu.Unlock()

	return nil
}

func (m *mem) Get(key string) (interface{}, error) {

	data, exist := m.storage[key]
	if !exist {
		return nil, errors.New("unknow key")
	}

	return data, nil
}

func (m *mem) Delete(key string) error {

	m.mu.Lock()
	{
		delete(m.storage, key)
	}
	m.mu.Unlock()

	return nil
}
