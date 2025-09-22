package stablemap

import (
	"slices"

	"github.com/vmihailenco/msgpack/v5"
)

type kvpair[K comparable, V any] struct {
	K K `msgpack:"k"`
	V V `msgpack:"v"`
}

func (sm *StableMap[K, V]) MarshalBinary() ([]byte, error) {
	if sm == nil {
		return nil, nil
	}
	pairs := make([]kvpair[K, V], 0, len(sm.index))
	for k, v := range sm.Entries() {
		kvp := kvpair[K, V]{k, v}
		pairs = append(pairs, kvp)
	}
	return msgpack.Marshal(pairs)
}

func (lm *LexicalStableMap[K, V]) MarshalBinary() ([]byte, error) {
	slices.Sort(lm.index)
	return lm.StableMap.MarshalBinary()
}

func (sm *StableMap[K, V]) UnmarshalBinary(p []byte) error {
	if p == nil {
		return nil
	}
	var pairs []kvpair[K, V]
	err := msgpack.Unmarshal(p, &pairs)
	if err != nil {
		return err
	}
	for _, pair := range pairs {
		k := pair.K
		v := pair.V
		sm.Set(k, v)
	}
	return nil
}
