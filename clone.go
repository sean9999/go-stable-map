package stablemap

// Clone performs a safe clone
func (sm *StableMap[K, V]) Clone() *StableMap[K, V] {
	b := New[K, V]()
	for k, v := range sm.Entries() {
		b.Set(k, v)
	}
	return b
}
