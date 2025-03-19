package main

import "fmt"

func main() {
	ints := []int{1, 2, 3}

	bigInts := []int64{1, 2, 3, 4, 5}

	sum := Sums(ints)
	bigSum := Sums(bigInts)

	fmt.Printf("Sum ints: %d\n", sum)
	fmt.Printf("Sum bigs: %d", bigSum)
}

func Sums[T int64 | int](items []T) T {
	var sum T
	for _, v := range items {
		sum += v
	}

	return sum
}
