package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Printf("线程%d, sum为:%d\n", 1, 2)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go AsyncFunc(i)
	}
	wg.Wait()
}
func AsyncFunc(index int) {
	sum := 0
	for i := 0; i < 10000; i++ {
		sum += 1

	}
	fmt.Printf("线程%d, sum为:%d\n", index, sum)
	wg.Done()
}
