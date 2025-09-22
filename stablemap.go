package stablemap

import (
	"cmp"
	"iter"
	"slices"
	"sync"
)

// StableMap is a map whose keys are ordered, and whose operations are concurrency safe
// and which marshals itself into predictable, deterministic bytes.
type StableMap[K comparable, V any] struct {
	*sync.RWMutex
	m     map[K]V
	index []K
}

// LexicalStableMap is a StableMap that cares about lexical order rather than insertion order.
type LexicalStableMap[K cmp.Ordered, V any] struct {
	*StableMap[K, V]
}

func New[K comparable, V any]() *StableMap[K, V] {
	return &StableMap[K, V]{
		RWMutex: &sync.RWMutex{},
		m:       map[K]V{},
		index:   []K{},
	}
}

func NewLexical[K cmp.Ordered, V any]() *LexicalStableMap[K, V] {
	return &LexicalStableMap[K, V]{
		StableMap: New[K, V](),
	}
}

func (l *LexicalStableMap[K, V]) Entries() iter.Seq2[K, V] {
	slices.Sort(l.index)
	return l.StableMap.Entries()
}

func From[K comparable, V any](m map[K]V) *StableMap[K, V] {
	sm := New[K, V]()
	if m != nil {
		sm.Incorporate(m)
	}
	return sm
}

func LexicalFrom[K cmp.Ordered, V any](m map[K]V) *LexicalStableMap[K, V] {
	sm := NewLexical[K, V]()
	if m != nil {
		sm.Incorporate(m)
	}
	slices.Sort(sm.index)
	return sm
}
