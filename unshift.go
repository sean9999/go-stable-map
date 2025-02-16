package stablemap

func (sm *StableMap[K, V]) unshift(key K, val V) {
	i := sm.indexOf(key)
	sm.index = append([]K{key}, omit(i, sm.index)...)
	sm.set(key, val)
}

// place this pair at the head
func (sm *StableMap[K, V]) Unshift(key K, val V) {
	sm.Lock()
	defer sm.Unlock()
	sm.unshift(key, val)
}
