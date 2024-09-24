package solutions

import (
	"math/rand/v2"
)

func GenerateSequentialSlice(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = i
	}
	return slice
}

func GenerateIdenticalSlice(n int, value int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = value
	}
	return slice
}

func GenerateRandomSlice(n int, max int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = rand.IntN(max)
	}
	return slice
}
