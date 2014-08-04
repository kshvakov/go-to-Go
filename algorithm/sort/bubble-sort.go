package sort

func BubbleSort(slice []int) {

	swapped := true

	for swapped {

		swapped = false

		for i := 0; i < len(slice)-1; i++ {

			j := i + 1

			if slice[i] > slice[j] {

				swapped = true

				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
}
