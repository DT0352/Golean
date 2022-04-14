package main

import (
	"fmt"
	"sync"
)

var m sync.Mutex
var set = make(map[int]bool, 0)

func printOnce(num int) {
	m.Lock()
	if _, exist := set[num]; !exist {
		fmt.Println(num)
	}
	set[num] = true
	m.Unlock()
}

//func main() {
//	for i := 0; i < 10; i++ {
//		go printOnce(100)
//	}
//	time.Sleep(time.Second)
//}
func main() {
	a := 1
	a -= 1 + 2
	fmt.Println(a)
}
