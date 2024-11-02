package stablemap

import (
	"github.com/vmihailenco/msgpack/v5"
)

type kvpair[K comparable, V any] struct {
	K K `msgpack:"k"`
	V V `msgpack:"v"`
}

func (sm *StableMap[K, V]) MarshalBinary() ([]byte, error) {

	pairs := make([]kvpair[K, V], 0, len(sm.index))
	for k, v := range sm.Entries() {
		kvp := kvpair[K, V]{k, v}
		pairs = append(pairs, kvp)
	}
	return msgpack.Marshal(pairs)
}

func (sm *StableMap[K, V]) UnmarshalBinary(p []byte) error {

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
