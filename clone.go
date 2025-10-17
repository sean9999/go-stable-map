package stablemap

// Clone performs a safe clone
func (sm *Map[K, V]) Clone() *Map[K, V] {
	b := New[K, V]()
	for k, v := range sm.Entries() {
		b.Set(k, v)
	}
	return b
}
