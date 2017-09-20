package rwmap

import (
	"math/rand"
	"sync"
	"time"
)

type Map interface {
	Load(key interface{}) (value interface{}, ok bool)
	Store(key, value interface{})
	LoadOrStore(key, value interface{}) (actual interface{}, loaded bool)
	Delete(key interface{})
	Range(f func(key, value interface{}) bool)
	Len() int
}

type rwMap struct {
	data map[interface{}]interface{}
	rwmu sync.RWMutex
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
func RwMap(capacity ...int) Map {
	var cap int
	if len(capacity) > 0 {
		cap = capacity[0]
	}
	return &rwMap{
		data: make(map[interface{}]interface{}, cap),
	}
}
func (m *rwMap) Load(key interface{}) (value interface{}, ok bool) {
	m.rwmu.RLock()
	value, ok = m.data[key]
	m.rwmu.RUnlock()
	return value, ok
}
func (m *rwMap) Store(key, value interface{}) {
	m.rwmu.Lock()
	m.data[key] = value
	m.rwmu.Unlock()
}
func (m *rwMap) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool) {
	m.rwmu.Lock()
	actual, loaded = m.data[key]
	m.data[key] = value
	if !loaded {
		actual = value
	}
	m.rwmu.Unlock()
	return actual, loaded
}
func (m *rwMap) Delete(key interface{}) {
	m.rwmu.Lock()
	delete(m.data, key)
	m.rwmu.Unlock()
}
func (m *rwMap) Range(f func(key, value interface{}) bool) {
	m.rwmu.RLock()
	defer m.rwmu.RUnlock()
	for k, v := range m.data {
		if !f(k, v) {
			break
		}
	}
}
func (m *rwMap) Len() int {
	m.rwmu.RLock()
	defer m.rwmu.RUnlock()
	return len(m.data)
}
