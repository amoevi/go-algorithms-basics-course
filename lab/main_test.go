package main

import (
	"lab/utils"
	"slices"
	"testing"
)

func TestGenerateSequentialSlice(t *testing.T) {
	// Helper function to test slices
	checkIsSequential := func(size int, expected []int) {
		generated := utils.GenerateSequentialSlice(size)
		if !slices.Equal(generated, expected) {
			t.Errorf("\n\texpected\t%v\n\tgenerated\t%v", expected, generated)
		}
	}

	expected1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	checkIsSequential(10, expected1)

	expected2 := []int{}
	checkIsSequential(0, expected2)

	expected3 := []int{0}
	checkIsSequential(1, expected3)
}

func TestGenerateIdenticalSlice(t *testing.T) {
	// Helper function to test slices
	checkIsIdentical := func(size int, value int, expected []int) {
		generated := utils.GenerateIdenticalSlice(size, value)
		if !slices.Equal(generated, expected) {
			t.Errorf("\n\texpected\t%v\n\tgenerated\t%v", expected, generated)
		}
	}

	expected1 := []int{7, 7, 7, 7, 7}
	checkIsIdentical(5, 7, expected1)

	expected2 := []int{}
	checkIsIdentical(0, 4, expected2)
}

func TestGenerateRandomSlice(t *testing.T) {
	n := 10
	max := 100
	generated := utils.GenerateRandomSlice(n, max)

	if len(generated) != n {
		t.Errorf("Expected slice length of %d, but got %d", n, len(generated))
	}

	for i, value := range generated {
		if value < 0 || value >= max {
			t.Errorf("Value out of range at index %d: got %d, but expected between 0 and %d", i, value, max-1)
		}
	}

	generated2 := utils.GenerateRandomSlice(n, max)
	if slices.Equal(generated, generated2) {
		t.Errorf("Expected two different slices but got the same: %v and %v", generated, generated2)
	}

	for i, value := range generated2 {
		if value < 0 || value >= max {
			t.Errorf("Value out of range at index %d: got %d, but expected between 0 and %d", i, value, max-1)
		}
	}
}

func TestReverseSlice(t *testing.T) {
	// Helper function to test slices
	checkReverse := func(slice []int, expected []int) {
		utils.ReverseSlice(slice)
		if !slices.Equal(slice, expected) {
			t.Errorf("\n\texpected\t%v\n\tbut got \t%v", expected, slice)
		}
	}

	// Test with a slice of multiple integers
	slice1 := []int{1, 2, 3, 4, 5}
	expected1 := []int{5, 4, 3, 2, 1}
	checkReverse(slice1, expected1)

	// Test with an empty slice
	slice2 := []int{}
	expected2 := []int{}
	checkReverse(slice2, expected2)

	// Test with a slice containing one element
	slice3 := []int{42}
	expected3 := []int{42}
	checkReverse(slice3, expected3)

	// Test with a slice containing two elements
	slice4 := []int{10, 20}
	expected4 := []int{20, 10}
	checkReverse(slice4, expected4)

	// Test with a slice containing repeated elements
	slice5 := []int{1, 2, 2, 3, 1}
	expected5 := []int{1, 3, 2, 2, 1}
	checkReverse(slice5, expected5)
}

func TestIsSorted(t *testing.T) {
	// Helper function to test slices
	checkIsSorted := func(slice []int, expected bool) {
		result := utils.IsSorted(slice)
		if result != expected {
			t.Errorf("Expected IsSorted to return %v for slice %v, but got %v.", expected, slice, result)
		}
	}

	// Test sorted slices
	slice1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	checkIsSorted(slice1, true)

	slice2 := []int{0, 0, 0, 0, 0, 0}
	checkIsSorted(slice2, true)

	slice3 := []int{}
	checkIsSorted(slice3, true)

	// Test unsorted slices
	slice4 := []int{1, 3, 2, 5, 4}
	checkIsSorted(slice4, false)

	slice5 := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	checkIsSorted(slice5, false)

	slice6 := []int{5, 1, 2, 3, 0}
	checkIsSorted(slice6, false)
}
