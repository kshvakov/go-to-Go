package sort

func OddEvenSort(slice []int) {

	sorted := false

	for !sorted {

		sorted = true

		for i := 1; i < len(slice)-1; i += 2 {

			if slice[i] > slice[i+1] {

				sorted = false

				slice[i], slice[i+1] = slice[i+1], slice[i]
			}
		}

		for i := 0; i < len(slice)-1; i += 2 {

			if slice[i] > slice[i+1] {

				sorted = false

				slice[i], slice[i+1] = slice[i+1], slice[i]
			}
		}
	}
}
