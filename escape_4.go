package main

import "fmt"

func series() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}

}
func main() {
	f := series()
	for i := 0; i < 10; i++ {
		fmt.Printf("Series:%d\n", f())
	}

}
