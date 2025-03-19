package main

import "fmt"

var cache = map[int]int{
	0: 0,
	1: 1,
	2: 1,
}

var called = 0

func main() {
	fib := Fib(100)

	fmt.Printf("result: %d\nCalled: %d", fib, called)
}

func Fib(depth int) int {
	d, found := cache[depth]

	if found {
		called += 1
		return d
	}

	var1 := Fib(depth - 1)
	var2 := Fib(depth - 2)

	cache[depth-1] = var1
	cache[depth-2] = var2

	return var1 + var2
}
