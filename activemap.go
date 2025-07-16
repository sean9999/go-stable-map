package stablemap

// A Result is emitted whenever an ActiveMap mutates
type Result[K comparable, V any] struct {
	Action string
	Key    K
	OldVal V
	NewVal V
	Msg    string
}

// ActiveMap is a StableMap that emits events when it mutates.
type ActiveMap[K comparable, V any] struct {
	*StableMap[K, V]
	events chan Result[K, V]
}

// NewActiveMap instantiates a new ActiveMap
func NewActiveMap[K comparable, V any]() *ActiveMap[K, V] {
	sm := New[K, V]()
	return &ActiveMap[K, V]{
		StableMap: sm,
	}
}

func (am *ActiveMap[K, V]) Set(k K, v V) error {
	am.Lock()
	defer am.Unlock()
	oldVal, _ := am.StableMap.get(k)
	err := am.StableMap.set(k, v)
	if err != nil {
		return err
	}
	res := Result[K, V]{
		Action: "set",
		Key:    k,
		OldVal: oldVal,
		NewVal: v,
	}
	if am.events != nil {
		am.events <- res
	}
	return nil
}

func (am *ActiveMap[K, V]) Delete(k K) error {
	am.Lock()
	defer am.Unlock()
	var zeroVal V
	oldVal, _ := am.StableMap.get(k)
	err := am.StableMap.delete(k)
	if err != nil {
		return err
	}
	res := Result[K, V]{
		Action: "delete",
		Key:    k,
		OldVal: oldVal,
		NewVal: zeroVal,
	}
	if am.events != nil {
		am.events <- res
	}
	return nil
}

func (am *ActiveMap[K, V]) Events() <-chan Result[K, V] {
	if am.events == nil {
		am.events = make(chan Result[K, V])
	}
	return am.events
}
