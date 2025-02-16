package stablemap

import (
	"fmt"
	"slices"
)

// delete the element at a specific index, or panic
func (sm *StableMap[K, V]) DeleteAt(i int) {
	k := sm.index[i]
	delete(sm.m, k)
	sm.removeElementAt(i)
}

func (sm *StableMap[K, V]) delete(k K) error {
	i := sm.indexOf(k)
	if i < 0 {
		return fmt.Errorf("key %v did not exist", k)
	}
	delete(sm.m, k)
	if i > -1 {
		sm.removeElementAt(i)
	}
	return nil
}

// delete an element by key
func (sm *StableMap[K, V]) Delete(k K) error {
	sm.Lock()
	defer sm.Unlock()
	return sm.delete(k)
}

func omit[T comparable](i int, things []T) []T {
	if i < 0 {
		return things
	}
	return append(things[:i], things[i+1:]...)
}

// remove element at the specified index, or panic
func (sm *StableMap[K, V]) removeElementAt(i int) {
	sm.index = slices.Delete(sm.index, i, i+1)
}
