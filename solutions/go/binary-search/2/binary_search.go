package binarysearch

func SearchInts(list []int, key int) int {
	left := 0
	right := len(list) - 1

	for left <= right {
		midpoint := (left + right) / 2

		if list[midpoint] > key {
			right = midpoint - 1
		} else if list[midpoint] < key {
			left = midpoint + 1
		} else {
			return midpoint
		}
	}

	return -1
}
