package stablemap

func (sm *StableMap[K, V]) AsMap() map[K]V {
	if sm == nil {
		return nil
	}
	m := make(map[K]V, sm.Length())
	for k, v := range sm.Entries() {
		m[k] = v
	}
	return m
}
