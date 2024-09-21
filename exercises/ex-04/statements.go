package main

import "fmt"

func main() {
	// You can modify the number variable to test other values.
	number := 1948764641

	/*
		- If the number is 1 or less, it's not prime.
		- We check each number from 2 up to the number itself :
			- If any of these numbers divide it, the number isn’t prime.
			- If none do, the number is prime.
	*/
	var isPrime bool

	// TODO

	if isPrime {
		fmt.Println(number, "is a prime number")
	} else {
		fmt.Println(number, "is not a prime number")
	}
}
