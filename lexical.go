package stablemap

import (
	"cmp"
	"iter"
	"slices"
)

// LexicalMap is a Map that cares about lexical order rather than insertion order.
type LexicalMap[K cmp.Ordered, V any] struct {
	*Map[K, V]
}

func LexicalFrom[K cmp.Ordered, V any](m map[K]V) *LexicalMap[K, V] {
	sm := NewLexicalMap[K, V]()
	if m != nil {
		sm.Incorporate(m)
	}
	slices.Sort(sm.index)
	return sm
}

func NewLexicalMap[K cmp.Ordered, V any]() *LexicalMap[K, V] {
	return &LexicalMap[K, V]{
		Map: New[K, V](),
	}
}

func (lm *LexicalMap[K, V]) Entries() iter.Seq2[K, V] {
	slices.Sort(lm.index)
	return lm.Map.Entries()
}
