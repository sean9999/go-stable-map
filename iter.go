package stablemap

import (
	"cmp"
	"iter"
	"slices"
)

// Entries provides a stable "range over" iteration, preserving insertion order.
func (sm *StableMap[K, V]) Entries() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for _, k := range sm.index {
			v, _ := sm.get(k)
			if !yield(k, v) {
				return
			}
		}
	}
}

func LexicalRange[K cmp.Ordered, V any](sm StableMap[K, V]) iter.Seq2[K, V] {
	slices.Sort(sm.index)
	return func(yield func(K, V) bool) {
		for _, k := range sm.index {
			v, _ := sm.get(k)
			if !yield(k, v) {
				return
			}
		}
	}
}
