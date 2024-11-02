# Stable Map

Stable Map is a map that provides ordered, deterministic iteration of key-value pairs. It's concurrency safe. Insertion order of key-value pairs is preserved. It marshals and unmarshals itself into deterministic bytes, using [MessagePack](https://msgpack.org/index.html).

Stable Map is especially useful for equality comparison.

## Getting Started

```go
package main

import (
	"fmt"

	smap "github.com/sean9999/go-stable-map"
)

func main() {

	m := smap.New[string, string]()

	m.Set("foo", "bar")
	m.Set("bing", "bat")

	//	dump the map 10 times. See that the order is always the same
	for range 10 {
		for k, v := range m.Entries() {
			fmt.Printf("%s:\t%s\n", k, v)
		}
		
		//	see that the binary representation is always the same
		bin, _ := m.MarshalBinary()
		fmt.Printf("hex:\t%x\n\n", bin)
	}

}
```
