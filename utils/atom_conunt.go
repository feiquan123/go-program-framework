package utils

import "sync"

// Counter : is a count
type Counter struct {
	Lock  *sync.RWMutex
	Count int64
}

// NewCounter : create a Counter
func NewCounter() *Counter {
	return &Counter{
		Lock:  new(sync.RWMutex),
		Count: 0,
	}
}

// Add : Counter add n
func (c *Counter) Add(n int64) {
	c.Lock.Lock()
	defer c.Lock.Unlock()

	c.Count += n
}

// Dec : Counter dec n
func (c *Counter) Dec(n int64) {
	c.Lock.Lock()
	defer c.Lock.Unlock()

	c.Count -= n
}

// Get : Get current num from Counter
func (c *Counter) Get() int64 {
	c.Lock.RLock()
	defer c.Lock.RUnlock()

	return c.Count
}

// Reset : reset Counter to zero
func (c *Counter) Reset() {
	c.Lock.Lock()
	defer c.Lock.Unlock()

	c.Count = 0
}
