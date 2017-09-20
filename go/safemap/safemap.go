package safemap

// ref: https://gist.github.com/fatih/6206844

import (
	"sync"
)

// SafeMap is concurrent security map
type SafeMap struct {
	lock *sync.RWMutex
	sm   map[interface{}]interface{}
}

// NewSafeMap get a new concurrent security map
func NewSafeMap() *SafeMap {
	return &SafeMap{
		lock: new(sync.RWMutex),
		sm:   make(map[interface{}]interface{}),
	}
}

// Get used to get a value by key
func (m *SafeMap) Get(k interface{}) interface{} {
	m.lock.RLock()
	defer m.lock.RUnlock()
	if val, ok := m.sm[k]; ok {
		return val
	}
	return nil
}

// Set used to set value with key
func (m *SafeMap) Set(k interface{}, v interface{}) bool {
	m.lock.RLock()
	defer m.lock.RUnlock()
	if val, ok := m.sm[k]; !ok {
		m.sm[k] = v
	} else if val != v {
		m.sm[k] = v
	} else {
		return false
	}
	return true
}

// Len returns the number of items in a set.
func (m *SafeMap) Len() int {
	return len(m.sm)
}

// Clear removes all items from the set
func (m *SafeMap) Clear() {
	m.lock.RLock()
	defer m.lock.RUnlock()
	m.sm = make(map[interface{}]interface{})
}

// IsExists determine whether k exists
func (m *SafeMap) IsExists(k interface{}) bool {
	m.lock.RLock()
	defer m.lock.RUnlock()
	if _, ok := m.sm[k]; !ok {
		return false
	}
	return true
}

// Delete used to delete a key
func (m *SafeMap) Delete(k interface{}) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	delete(m.sm, k)
}

// IsEmpty checks for emptiness
func (m *SafeMap) IsEmpty() bool {
	if len(m.sm) == 0 {
		return true
	}
	return false
}

// Set returns a slice of all items
func (m *SafeMap) List() []interface{} {
	m.lock.RLock()
	defer m.lock.RUnlock()
	list := make([]interface{}, 0)
	for item := range m.sm {
		list = append(list, item)
	}
	return list
}
