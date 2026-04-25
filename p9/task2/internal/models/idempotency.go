package models

import "sync"

type CachedResponse struct {
	StatusCode int
	Body       []byte
	Completed  bool
}

type InMemory struct {
	mu   sync.Mutex
	data map[string]*CachedResponse
}

func NewInMemory() *InMemory {
	return &InMemory{data: make(map[string]*CachedResponse)}
}

func (m *InMemory) Get(key string) (*CachedResponse, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	resp, exists := m.data[key]
	return resp, exists
}

func (m *InMemory) StartProcessing(key string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.data[key]; exists {
		return false
	}

	m.data[key] = &CachedResponse{}

	return true
}

func (m *InMemory) Finish(key string, status int, body []byte) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if resp, exists := m.data[key]; exists {
		resp.StatusCode = status
		resp.Body = body
		resp.Completed = true
	} else {
		m.data[key] = &CachedResponse{StatusCode: status, Body: body, Completed: true}
	}
}
