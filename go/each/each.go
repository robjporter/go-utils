package each

import (
	"reflect"
	"sync"
)

func init() {
	MakeEach(&Each)
	MakeEach(&EachInt)
	// MakeEach(&EachString)
	MakeEach(&EachStringInt)
	MakeEachP(&EachP)
}

var Each func(fn interface{}, slice_or_map interface{})
var EachP func(fn interface{}, slice_or_map interface{})
var EachInt func(func(value, i int), []int)
var EachStringInt func(func(value int, key string), map[string]int)

func Maker(fn interface{}, impl func(args []reflect.Value) (results []reflect.Value)) {
	fnV := reflect.ValueOf(fn).Elem()
	fnI := reflect.MakeFunc(fnV.Type(), impl)
	fnV.Set(fnI)
}

func interfaceToValue(v reflect.Value) reflect.Value {
	if v.Kind() == reflect.Interface {
		return reflect.ValueOf(v.Interface())
	}
	return v
}

func extractArgs(values []reflect.Value) (reflect.Value, reflect.Value) {
	fn := interfaceToValue(values[0])
	col := interfaceToValue(values[1])
	return fn, col
}

func MakeEach(fn interface{}) {
	Maker(fn, each)
}

func MakeEachP(fn interface{}) {
	Maker(fn, eachP)
}

func each(values []reflect.Value) []reflect.Value {
	fn, col := extractArgs(values)

	if col.Kind() == reflect.Map {
		eachMap(fn, col)
	}

	if col.Kind() == reflect.Slice {
		eachSlice(fn, col)
	}

	return nil
}

func eachSlice(fn, s reflect.Value) {
	for i := 0; i < s.Len(); i++ {
		v := s.Index(i)
		eachCall(fn, v, reflect.ValueOf(i))
	}
}

func eachMap(fn, m reflect.Value) {
	for _, k := range m.MapKeys() {
		v := m.MapIndex(k)
		eachCall(fn, v, k)
	}
}

func eachCall(fn, v, i reflect.Value) {
	args := []reflect.Value{v}
	if in := fn.Type().NumIn(); in == 2 {
		args = append(args, i)
	}
	fn.Call(args)
}

func eachP(values []reflect.Value) []reflect.Value {

	fn, col := extractArgs(values)

	if col.Kind() == reflect.Map {
		eachMapP(fn, col)
	}

	if col.Kind() == reflect.Slice {
		eachSliceP(fn, col)
	}

	return nil
}

func eachSliceP(fn, s reflect.Value) {
	var done sync.WaitGroup
	for i := 0; i < s.Len(); i++ {
		v := s.Index(i)
		done.Add(1)
		go func() {
			eachCall(fn, v, reflect.ValueOf(i))
			done.Done()
		}()
	}
	done.Wait()
}

func eachMapP(fn, m reflect.Value) {
	var done sync.WaitGroup
	keys := m.MapKeys()
	done.Add(len(keys))

	for _, k := range keys {
		v := m.MapIndex(k)
		go func(fn, v, k reflect.Value) {
			eachCall(fn, v, k)
			done.Done()
		}(fn, v, k)
	}
	done.Wait()
}

func refEach(slice []string, fn func(string)) {
	for i := 0; i < len(slice); i++ {
		fn(slice[i])
	}
}

func refPEach(slice []string, fn func(string)) {
	var done sync.WaitGroup

	for _, s := range slice {
		s := s
		done.Add(1)
		go func() {
			fn(s)
			done.Done()
		}()
	}

	done.Wait()
}
