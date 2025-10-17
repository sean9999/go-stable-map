package stablemap

func (sm *Map[K, V]) getAt(i int) V {
	key := sm.index[i]
	val := sm.m[key]
	return val
}

// get the element at a particular index, or panic
func (sm *Map[K, V]) GetAt(i int) V {
	sm.RLock()
	defer sm.RUnlock()
	return sm.getAt(i)
}
