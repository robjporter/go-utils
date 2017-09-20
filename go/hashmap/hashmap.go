package hashmap

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
	"sync"
)

// Key-value pairs for semantic logging
type HashMap struct {
	sync.RWMutex
	kv map[string]interface{}
}

func (hm *HashMap) Get(key string) interface{} {
	hm.RLock()
	defer hm.RUnlock()
	return hm.kv[key]
}

func (hm *HashMap) Set(key string, value interface{}) {
	hm.Lock()
	defer hm.Unlock()
	hm.kv[key] = value
}

func (hm *HashMap) SetFromMap(m map[string]interface{}) {
	hm.Lock()
	defer hm.Unlock()

	for key, value := range m {
		hm.kv[key] = value
	}
}

func (hm *HashMap) Contains(key string) bool {
	hm.Lock()
	defer hm.Unlock()
	_, exists := hm.kv[key]

	return exists
}

func NewHashMap() *HashMap {
	m := make(map[string]interface{})
	return &HashMap{kv: m}
}

// Sanitize the key for logging purposes
func encodeKey(k string) (key string) {
	// Keys may not have any spaces
	key = strings.Replace(k, " ", "_", -1)

	return key
}

// Encode the value of the map for certain supported types.
func encodeValue(i interface{}) (buf []byte, err error) {
	v := reflect.ValueOf(i)

	switch v.Kind() {
	case reflect.String:
		buf = append(buf, fmt.Sprintf("%q", i)...)
	case reflect.Array, reflect.Chan, reflect.Func, reflect.Interface,
		reflect.Map, reflect.Ptr, reflect.Struct:
		err = fmt.Errorf("Unable to encode %s value: type=%s value=%v",
			v.Kind(), reflect.TypeOf(i), i)
	default:
		buf = append(buf, fmt.Sprintf("%v", i)...)
	}

	return buf, err
}

func (hm *HashMap) String() string {
	var buf []byte
	var keys []string
	hm.RLock()
	defer hm.RUnlock()

	for k := range hm.kv {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		value, err := encodeValue(hm.kv[k])

		if err != nil {
			continue
		}

		key := encodeKey(k)
		buf = append(buf, fmt.Sprintf(" %s=%s", key, value)...)
	}

	// Remove leading space
	buf = buf[1:len(buf)]

	return string(buf)
}
