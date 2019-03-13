package main

import "fmt"

func fibonacci() func() int {
	a, b, cur := 0, 1, 0
	return func() int {
		a, b, cur = b, a+b, a
		return cur
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
