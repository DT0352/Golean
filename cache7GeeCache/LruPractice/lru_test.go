package lruPractice

import "testing"

type String string

func (d String) Len() int {
	return len(d)
}
func TestLRU_Get(t *testing.T) {

	lru := New(0, nil)
	lru.Add("k1", String("v1"))
}
