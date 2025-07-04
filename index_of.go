package stablemap

func (sm *StableMap[K, V]) indexOf(key K) int {
	for i, k := range sm.index {
		if k == key {
			return i
		}
	}

	return -1
}

// get the index value of a particular key
func (sm *StableMap[K, V]) IndexOf(key K) int {
	sm.RLock()
	defer sm.RUnlock()
	return sm.indexOf(key)
}
