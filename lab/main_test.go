package main

import (
	"lab/utils"
	"slices"
	"testing"
)

func TestGenerateSequentialSlice(t *testing.T) {
	generated := utils.GenerateSequentialSlice(10)
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	if !slices.Equal(generated, expected) {
		t.Fatalf("\n\texpected\t%v\n\tgenerated\t%v", expected, generated)
	}
}

func TestGenerateIdenticalSlice(t *testing.T) {
	generated := utils.GenerateIdenticalSlice(5, 7)
	expected := []int{7, 7, 7, 7, 7}
	if !slices.Equal(generated, expected) {
		t.Fatalf("\n\texpected\t%v\n\tgenerated\t%v", expected, generated)
	}
}

func TestGenerateRandomSlice(t *testing.T) {
	n := 10
	max := 100
	generated := utils.GenerateRandomSlice(n, max)

	if len(generated) != n {
		t.Fatalf("Expected slice length of %d, but got %d", n, len(generated))
	}

	for i, value := range generated {
		if value < 0 || value >= max {
			t.Fatalf("Value out of range at index %d: got %d, but expected between 0 and %d", i, value, max-1)
		}
	}

	generated2 := utils.GenerateRandomSlice(n, max)
	if slices.Equal(generated, generated2) {
		t.Fatalf("Expected two different slices but got the same: %v and %v", generated, generated2)
	}

	for i, value := range generated2 {
		if value < 0 || value >= max {
			t.Fatalf("Value out of range at index %d: got %d, but expected between 0 and %d", i, value, max-1)
		}
	}
}

func TestIsSorted(t *testing.T) {
	// Helper function to test sorted slices
	checkIsSorted := func(slice []int, expected bool) {
		result := utils.IsSorted(slice)
		if result != expected {
			t.Fatalf("Expected IsSorted to return %v for slice %v, but got %v.", expected, slice, result)
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
