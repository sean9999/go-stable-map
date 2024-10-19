package stablemap

import (
	"github.com/vmihailenco/msgpack/v5"
)

func (sm *StableMap[K, V]) MarshalBinary() ([]byte, error) {

	type kvpair struct {
		K K
		V V
	}
	pairs := make([]kvpair, 0, len(sm.index))
	for k, v := range sm.Entries() {
		kvp := kvpair{k, v}
		pairs = append(pairs, kvp)
	}
	return msgpack.Marshal(pairs)
}

func (sm *StableMap[K, V]) UnmarshalBinary(p []byte) error {

	type kvpair struct {
		K K
		V V
	}

	var pairs []kvpair

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
