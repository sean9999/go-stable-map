package stablemap

import "github.com/sean9999/pear"

func (sm *Map[K, V]) Length() int {
	sm.RLock()
	defer sm.RUnlock()
	if len(sm.index) != len(sm.m) {
		err := pear.Errorf("length mismatch. index is %d and map is %d", len(sm.index), len(sm.m))
		panic(err)
	}
	return len(sm.index)
}
