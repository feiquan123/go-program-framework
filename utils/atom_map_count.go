package utils

import "sync"

// MapCounter : map[string]Counter
type MapCounter struct {
	Lock     *sync.RWMutex
	Counters map[string]*Counter
}

// NewMapCounter : create a map counter
func NewMapCounter() *MapCounter {
	return &MapCounter{
		Lock:     new(sync.RWMutex),
		Counters: make(map[string]*Counter, 0),
	}
}

// Add : add key to map count
func (m *MapCounter) Add(key string, n int64) {
	m.Lock.Lock()
	defer m.Lock.Unlock()

	if mapCounter, ok := m.Counters[key]; ok {
		// update
		mapCounter.Add(n)
	} else {
		// add
		counter := NewCounter()
		counter.Add(n)
		m.Counters[key] = counter
	}
}

// Reset : reset map count
func (m *MapCounter) Reset() {
	m.Lock.Lock()
	defer m.Lock.Unlock()

	m.Counters = make(map[string]*Counter, 0)
}
