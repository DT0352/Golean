package lruPractice

import (
	"fmt"
	"testing"
)

type String string

func (d String) Len() int {
	return len(d)
}
func TestLRU_Get(t *testing.T) {

	lru := New(0, nil)
	lru.Add("k1", String("v1"))
	lru.Add("k2", String("v2"))
	lru.Add("k3", String("v3"))
	if ele, ok := lru.Get("k1"); ok {
		fmt.Println(ele)
	}

}
