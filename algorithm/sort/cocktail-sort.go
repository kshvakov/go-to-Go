package sort

func CocktailSort(slice []int) {

	left := 0
	right := len(slice) - 1

	for left <= right {

		for i := right; i > left; i-- {

			if slice[i] < slice[i-1] {

				slice[i], slice[i-1] = slice[i-1], slice[i]
			}
		}

		left++

		for i := left; i < right; i++ {

			if slice[i] > slice[i+1] {

				slice[i], slice[i+1] = slice[i+1], slice[i]
			}
		}

		right--
	}
}
