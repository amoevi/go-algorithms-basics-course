package main

import "fmt"

func slicesEqual(a, b []int) bool {
	var areSlicesEqual bool

	// TODO

	return areSlicesEqual
}

func main() {
	slice1 := []int{3, 1, 4, 5, 9}
	slice2 := []int{3, 1, 4, 5, 9}
	slice3 := []int{2, 7, 1}
	slice4 := []int{}
	slice5 := []int{3, 1, 4, 5, 9, 2}

	fmt.Println(slicesEqual(slice1, slice2))
	fmt.Println(slicesEqual(slice1, slice3))
	fmt.Println(slicesEqual(slice1, slice4))
	fmt.Println(slicesEqual(slice1, slice5))
}
