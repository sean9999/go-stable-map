package stablemap

import "errors"

// add or update an element
func (sm *StableMap[K, V]) set(key K, val V) error {
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

// add or update an element
func (sm *StableMap[K, V]) Set(key K, val V) error {
	sm.Lock()
	defer sm.Unlock()
	return sm.set(key, val)
}
