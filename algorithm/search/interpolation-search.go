package search

func InterpolationSearch(needle int, haystack []int) int {

	if len(haystack) == 0 {

		return -1
	}

	low := 0
	mid := 0
	high := len(haystack) - 1

	for haystack[high] != haystack[low] && haystack[low] <= needle && haystack[high] >= needle {

		mid = low + ((needle-haystack[low])*(high-low))/(haystack[high]-haystack[low])

		if haystack[mid] < needle {

			low = mid + 1

		} else if haystack[mid] > needle {

			high = mid - 1

		} else {

			return mid
		}
	}

	if haystack[low] == needle {

		return low
	}

	return -1
}
