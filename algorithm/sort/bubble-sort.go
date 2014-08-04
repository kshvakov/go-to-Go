package sort

func BubbleSort(slice []int) {

	swap := true

	for swap {

		swap = false

		for i := 0; i < len(slice)-1; i++ {

			j := i + 1

			if slice[i] > slice[j] {

				swap = true

				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
}
