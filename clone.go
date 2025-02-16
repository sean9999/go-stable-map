package stablemap

// safe clone
func (a *StableMap[K, V]) Clone() *StableMap[K, V] {
	b := New[K, V]()
	for k, v := range a.Entries() {
		b.Set(k, v)
	}
	return b
}
