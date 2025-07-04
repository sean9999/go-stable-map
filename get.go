package stablemap

func (sm *StableMap[K, V]) get(k K) (V, bool) {
	v, exists := sm.m[k]
	return v, exists
}

// get a value and a bolean indicating if there actually was something there
func (sm *StableMap[K, V]) Get(k K) (V, bool) {
	sm.RLock()
	defer sm.RUnlock()
	return sm.get(k)
}
