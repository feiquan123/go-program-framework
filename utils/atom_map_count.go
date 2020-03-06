package utils

import "sync"

// MapCounter : map[string]Counter
type MapCounter struct {
	Lock     *sync.RWMutex
	Counters map[string]int64
}

// NewMapCounter : create a map counter
func NewMapCounter() *MapCounter {
	return &MapCounter{
		Lock:     new(sync.RWMutex),
		Counters: make(map[string]int64, 0),
	}
}

// Add : add key to map count
func (m *MapCounter) Add(key string, n int64) {
	m.Lock.Lock()
	defer m.Lock.Unlock()

	if _, ok := m.Counters[key]; ok {
		// update
		m.Counters[key] += n
	} else {
		// add
		m.Counters[key] = n
	}
}

// Get : get count of one key
func (m *MapCounter) Get(key string) (n int64, ok bool) {
	m.Lock.RLock()
	defer m.Lock.RUnlock()

	n, ok = m.Counters[key]
	return
}

// Reset : reset map count
func (m *MapCounter) Reset() {
	m.Lock.Lock()
	defer m.Lock.Unlock()

	m.Counters = make(map[string]int64, 0)
}
