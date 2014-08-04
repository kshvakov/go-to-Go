package search

func BinarySearch(needle int, haystack []int) int {

	if len(haystack) == 0 {

		return -1
	}

	min := 0
	max := len(haystack) - 1

	if haystack[min] > needle || haystack[max] < needle {

		return -1
	}

	for min < max {

		mid := min + (max-min)/2

		if haystack[mid] < needle {

			min = mid + 1

		} else {

			max = mid
		}
	}

	if haystack[min] == needle {

		return min
	}

	return -1
}
