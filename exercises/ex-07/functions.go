package main

import "fmt"

func isLetter(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

func isDigit(char rune) bool {
	return char >= '0' && char <= '9'
}

// TODO
// Create the function countLettersAndDigits that takes a string
// and returns two integers: the number of letters and the number of digits.

// BONUS TODO
// Create the function isVowel that takes a rune as input
// and returns true if the rune is a vowel (a, e, i, o, u - both upper and lowercase).

// BONUS TODO
// Create the function isConsonant that takes a rune as input
// and returns true if the rune is a consonant (i.e., a letter but not a vowel).

// BONUS TODO
// Create the function countVowelsAndConsonants that takes a string
// and returns two integers: the number of vowels and the number of consonants.

func main() {
	// TODO
	// Call the function countLettersAndDigits with a string argument
	// Store the returned result in two variables and print them.

	// BONUS TODO
	// Call the function countVowelsAndConsonants with a string argument
	// Store the returned result in two variables and print them.

	fmt.Println()
}
