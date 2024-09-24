package utils

// TODO
// Create the function GenerateSequentialSlice that takes an integer n
// and returns a slice containing numbers from 0 to n-1
func GenerateSequentialSlice(n int) []int {
	slice := make([]int, n)
	return slice
}

// TODO
// Create the function GenerateIdenticalSlice that takes an integer n and an integer value
// and returns a slice of n integers where every element is equal to the given value

// TODO
// Create the function GenerateRandomSlice that takes an integer n and an integer max
// and returns a slice of n random integers.
//
// Note: In "math/rand/v2" package, rand.IntN(max) generates a random number between 0 and max-1.
// For example, rand.IntN(100) will generate a number between 0 and 99.
