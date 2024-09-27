package solutions

func IsSorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

func SelectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func MergeSort(arr []int) {
	// Base case: if the array has 1 or 0 elements, it's already sorted
	if len(arr) <= 1 {
		return
	}

	// Find the middle of the array
	mid := len(arr) / 2

	// Split the array into two halves
	left := arr[:mid]
	right := arr[mid:]

	// Recursively apply MergeSort to both halves
	MergeSort(left)
	MergeSort(right)

	// Merge the two sorted halves
	merge(arr, left, right)
}

// merge: merges two sorted subarrays into the original array
func merge(arr, left, right []int) {
	i, j, k := 0, 0, 0

	// Compare elements from both subarrays and insert into the main array
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
		k++
	}

	// Add remaining elements from left, if any
	for i < len(left) {
		arr[k] = left[i]
		i++
		k++
	}

	// Add remaining elements from right, if any
	for j < len(right) {
		arr[k] = right[j]
		j++
		k++
	}
}
