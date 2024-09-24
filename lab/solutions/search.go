package solutions

func LinearSearch(slice []int, target int) int {
	for i, v := range slice {
		if v == target {
			return i
		}
	}
	return -1
}

func BinarySearch(slice []int, target int) int {
	low, high := 0, len(slice)-1

	for low <= high {
		mid := (low + high) / 2
		if slice[mid] == target {
			return mid
		} else if slice[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}
