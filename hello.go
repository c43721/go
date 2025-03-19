package main

import "fmt"

func main() {
	ints := []int64{1, 2, 3}

	sum := Sums(ints)

	fmt.Printf("Sum: %d", sum)
}

func Sums(items []int64) int64 {
	var sum int64
	for _, v := range items {
		sum += v
	}

	return sum
}
