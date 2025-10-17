package stablemap

// Import imports another Map
func (sm *Map[K, V]) Import(b *Map[K, V]) {
	if b == nil || sm == nil {
		return
	}
	for k, v := range b.Entries() {
		sm.Set(k, v)
	}
}

// Incorporate incorporates a map, merging it with existing entries
func (sm *Map[K, V]) Incorporate(m map[K]V) {
	if m == nil {
		return
	}
	for k, v := range m {
		sm.Set(k, v)
	}
}
