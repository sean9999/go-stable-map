package stablemap

import "iter"

// Entries() provides stable "range over" iteration
func (sm *StableMap[K, V]) Entries() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for _, k := range sm.index {
			v, _ := sm.Get(k)
			if !yield(k, v) {
				return
			}
		}
	}
}
