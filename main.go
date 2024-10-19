package stablemap

import (
	"fmt"
	"iter"
	"slices"
	"sync"
)

// StableMap is a map whose keys are ordered, and whose operations are concurrency safe
// and which marshals itself into predictable, deterministic bytes
type StableMap[K comparable, V any] struct {
	sync.RWMutex
	m     map[K]V
	index []K
}

// get the element at a particular index, or panic
func (sm *StableMap[K, V]) GetAt(i int) V {
	sm.RLock()
	defer sm.RUnlock()
	key := sm.index[i]
	val := sm.m[key]
	return val

}

// get the index value of a particular key
func (sm *StableMap[K, V]) IndexOf(key K) int {
	sm.RLock()
	defer sm.RUnlock()
	for i, k := range sm.index {
		if k == key {
			return i
		}
	}

	return -1
}

// delete the element at a specific index, or panic
func (sm *StableMap[K, V]) DeleteAt(i int) {
	k := sm.index[i]
	delete(sm.m, k)
	sm.removeElementAtIndex(i)
}

// remove element at the specified index, or panic
func (sm *StableMap[K, V]) removeElementAtIndex(i int) {
	sm.index = slices.Delete(sm.index, i, i+1)
}

// delete an element
func (sm *StableMap[K, V]) Delete(k K) error {
	i := sm.IndexOf(k)
	if i < 0 {
		return fmt.Errorf("key %v did not exist", k)
	}
	sm.Lock()
	defer sm.Unlock()
	delete(sm.m, k)
	if i > -1 {
		sm.removeElementAtIndex(i)
	}
	return nil
}

// add or update an element
func (sm *StableMap[K, V]) Set(key K, val V) {
	sm.Lock()
	defer sm.Unlock()
	_, exists := sm.m[key]
	sm.m[key] = val
	if !exists {
		sm.index = append(sm.index, key)
	}
}

// get a value and a bolean indicating if there actually was something there
func (sm *StableMap[K, V]) Get(k K) (V, bool) {
	sm.RLock()
	defer sm.RUnlock()
	v, exists := sm.m[k]
	return v, exists
}

func (sm *StableMap[K, V]) Length() int {
	sm.RLock()
	defer sm.RUnlock()
	if len(sm.index) != len(sm.m) {
		panic("length mismatch")
	}
	return len(sm.index)
}

// Entries provides stable range over
func (sm *StableMap[K, V]) Entries() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for _, k := range sm.index {
			v, _ := sm.Get(k)
			if !yield(k, v) {
				return
			}
		}
	}
}

func New[K comparable, V any]() *StableMap[K, V] {
	return &StableMap[K, V]{
		m:     map[K]V{},
		index: []K{},
	}
}
