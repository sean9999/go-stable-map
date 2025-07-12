package stablemap

import (
	"sync"
)

// StableMap is a map whose keys are ordered, and whose operations are concurrency safe
// and which marshals itself into predictable, deterministic bytes.
type StableMap[K comparable, V any] struct {
	*sync.RWMutex
	m     map[K]V
	index []K
}

func New[K comparable, V any]() *StableMap[K, V] {
	return &StableMap[K, V]{
		RWMutex: &sync.RWMutex{},
		m:       map[K]V{},
		index:   []K{},
	}
}

func From[K comparable, V any](m map[K]V) *StableMap[K, V] {
	sm := New[K, V]()
	sm.Incorporate(m)
	return sm
}
