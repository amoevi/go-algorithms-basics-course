package solutions

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
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

func GenerateUserSlice() []int {
	reader := bufio.NewReader(os.Stdin)

	// Ask the user to input all the elements separated by spaces
	fmt.Print("Enter the elements separated by space: ")

	// Read the entire line, including the newline
	input, _ := reader.ReadString('\n')

	// Remove any trailing newline or spaces
	input = strings.TrimSpace(input)

	// Split the input string by spaces
	strElements := strings.Split(input, " ")

	// Create a slice of integers with a length equal to the number of string elements
	slice := make([]int, len(strElements))

	// Convert each string element to an integer and store it in the slice
	for i, str := range strElements {
		for {
			num, err := strconv.Atoi(str)
			if err != nil {
				// If there's an error, prompt the user to enter a valid integer
				fmt.Printf("Invalid input: %s is not an integer. Please enter a valid integer for element %d: ", str, i)
				newInput, _ := reader.ReadString('\n') // Use bufio again to capture the correct input
				str = strings.TrimSpace(newInput)      // Trim any extra spaces or newlines
			} else {
				slice[i] = num
				break // Exit the loop when a valid integer is entered
			}
		}
	}

	return slice
}
