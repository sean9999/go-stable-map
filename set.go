package stablemap

import "errors"

// add or update an element
func (sm *Map[K, V]) set(key K, val V) error {
	var zeroVal K
	if key == zeroVal {
		return errors.New("zero value key not allowed")
	}
	_, exists := sm.m[key]
	sm.m[key] = val
	if !exists {
		sm.index = append(sm.index, key)
	}
	return nil
}

// Set adds or updates an element in a thread-safe manner
func (sm *Map[K, V]) Set(key K, val V) error {
	sm.Lock()
	defer sm.Unlock()
	return sm.set(key, val)
}
