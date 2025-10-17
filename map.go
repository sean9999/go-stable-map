package stablemap

import (
	"sync"
)

// Map is a map whose keys are ordered, and whose operations are concurrency safe
// and which marshals itself into predictable, deterministic bytes.
type Map[K comparable, V any] struct {
	*sync.RWMutex
	m     map[K]V
	index []K
}

func New[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{
		RWMutex: &sync.RWMutex{},
		m:       map[K]V{},
		index:   []K{},
	}
}

func From[K comparable, V any](m map[K]V) *Map[K, V] {
	sm := New[K, V]()
	if m != nil {
		sm.Incorporate(m)
	}
	return sm
}
