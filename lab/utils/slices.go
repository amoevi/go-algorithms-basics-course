package utils

import (
	"bufio"
	"fmt"
	"os"
)

// TODO
// Create the function GenerateSequentialSlice that takes an integer n
// and returns a slice containing numbers from 0 to n-1
func GenerateSequentialSlice(n int) []int {
	var slice []int
	return slice
}

// TODO
// Create the function GenerateIdenticalSlice that takes an integer n and an integer value
// and returns a slice of n integers where every element is equal to the given value
func GenerateIdenticalSlice(n int, value int) []int {
	var slice []int
	return slice
}

// TODO
// Create the function GenerateRandomSlice that takes an integer n and an integer max
// and returns a slice of n random integers.
//
// Note: In "math/rand/v2" package, rand.IntN(max) generates a random number between 0 and max-1.
// For example, rand.IntN(100) will generate a number between 0 and 99.
func GenerateRandomSlice(n int, max int) []int {
	var slice []int
	return slice
}

// BONUS TODO
// Create the function GenerateUserSlice that asks the user
// to input all the elements in a single line, separated by spaces.
//
// HINTS:
// You may need to import the following packages:
// - bufio (for reading user input as a full line)
// - fmt (for printing output and handling input)
// - strings (for splitting the input string into individual values)
// - strconv (for converting strings to integers)
func GenerateUserSlice() []int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter some text: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("You entered:", input)

	var slice []int
	return slice
}
