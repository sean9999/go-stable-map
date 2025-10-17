package stablemap

func (sm *Map[K, V]) AsMap() map[K]V {
	// sm.Lock()
	// defer sm.Unlock()
	if sm == nil {
		return nil
	}
	m := make(map[K]V, sm.Length())
	for k, v := range sm.Entries() {
		m[k] = v
	}
	return m
}
