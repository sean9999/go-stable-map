package stablemap_test

import (
	"testing"

	stablemap "github.com/sean9999/go-stable-map"
	"github.com/stretchr/testify/assert"
)

func TestStableMap_Set(t *testing.T) {
	sm := stablemap.New[string, int]()
	sm.Set("one", 1)
	val, exists := sm.Get("one")
	assert.True(t, exists)
	assert.Equal(t, 1, val)
	assert.Equal(t, 1, sm.Length())

	sm.Set("two", 2)
	val, exists = sm.Get("two")
	assert.True(t, exists)
	assert.Equal(t, 2, val)
	assert.Equal(t, 2, sm.Length())

	sm.Set("two", 22)
	val, exists = sm.Get("two")
	assert.True(t, exists)
	assert.Equal(t, 22, val)
	assert.Equal(t, 2, sm.Length())

	sm.Delete("two")
	val, exists = sm.Get("two")
	assert.False(t, exists)
	assert.Equal(t, 0, val)
	assert.Equal(t, 1, sm.Length())

	sm.Set("two", 2)
	val, exists = sm.Get("two")
	assert.True(t, exists)
	assert.Equal(t, 2, val)
	assert.Equal(t, 2, sm.Length())

	sm.Set("three", 3)
	sm.Set("four", 4)
	sm.Set("five", 5)
	sm.Set("six", 6)
	sm.Set("seven", 7)
	sm.Set("eight", 8)

	lastInteger := 0
	for k, v := range sm.Entries() {
		if v <= lastInteger {
			t.Errorf("expected a value higher than %d but got %d with key %q", lastInteger, v, k)
			break
		}
		lastInteger = v
	}

	sixth := sm.GetAt(5)
	assert.Equal(t, 6, sixth)

	bytes1, err := sm.MarshalBinary()
	assert.NoError(t, err)

	sm2 := stablemap.New[string, int]()

	err = sm2.UnmarshalBinary(bytes1)
	assert.NoError(t, err)

	bytes2, err := sm2.MarshalBinary()
	assert.NoError(t, err)

	assert.Equal(t, bytes1, bytes2)

}

func TestLexicalRange(t *testing.T) {
	sm := stablemap.New[string, int]()
	sm.Set("fig", 6)
	sm.Set("apple", 1)
	sm.Set("cherry", 3)
	sm.Set("date", 4)
	sm.Set("elderberry", 5)
	sm.Set("banana", 2)
	i := 0
	ovals := []string{
		"apple",
		"banana",
		"cherry",
		"date",
		"elderberry",
		"fig",
	}
	for k, _ := range stablemap.LexicalRange[string, int](*sm) {
		j := sm.IndexOf(k)
		assert.Equal(t, i, j)
		assert.Equal(t, ovals[i], k)
		i++
	}
}

func TestLexicalStableMap(t *testing.T) {
	sm := stablemap.NewLexical[string, int]()
	sm.Set("fig", 6)
	sm.Set("apple", 1)
	sm.Set("cherry", 3)
	sm.Set("date", 4)
	sm.Set("elderberry", 5)
	sm.Set("banana", 2)
	i := 0
	ovals := []string{
		"apple",
		"banana",
		"cherry",
		"date",
		"elderberry",
		"fig",
	}
	for k, _ := range sm.Entries() {
		j := sm.IndexOf(k)
		assert.Equal(t, i, j)
		assert.Equal(t, ovals[i], k)
		i++
	}
}
