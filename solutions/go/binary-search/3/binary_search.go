package binarysearch

func SearchInts(list []int, key int) int {
	left := 0
	right := len(list) - 1

	for left <= right {
		midpoint := (left + right) / 2

		switch {
		case list[midpoint] > key:
			right = midpoint - 1
		case list[midpoint] < key:
			left = midpoint + 1
		default:
			return midpoint
		}
	}

	return -1
}
