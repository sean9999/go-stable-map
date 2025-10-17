package stablemap

import (
	"encoding/json"
)

type kvset [][2]any

func (sm Map[K, V]) MarshalJSON() ([]byte, error) {
	var zeroValue K
	omap := make(kvset, 0)
	for k, v := range sm.Entries() {
		if k != zeroValue {
			tuple := [2]any{
				k, v,
			}
			omap = append(omap, tuple)
		}
	}
	return json.Marshal(omap)
}

func (sm *Map[K, V]) UnmarshalJSON(b []byte) error {
	var omap kvset
	err := json.Unmarshal(b, &omap)
	if err != nil {
		return err
	}
	for _, kv := range omap {
		k, ok := kv[0].(K)
		if ok {
			sm.Set(k, kv[1].(V))
		}
	}
	return nil
}
