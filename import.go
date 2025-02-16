package stablemap

// Import imports another StableMap
func (a *StableMap[K, V]) Import(b *StableMap[K, V]) {
	if b == nil || a == nil {
		return
	}
	for k, v := range b.Entries() {
		a.Set(k, v)
	}
}

// Incorporate incoporates a map, merging it with existing entries
func (sm *StableMap[K, V]) Incorporate(m map[K]V) {
	if m == nil {
		return
	}
	for k, v := range m {
		sm.Set(k, v)
	}
}
