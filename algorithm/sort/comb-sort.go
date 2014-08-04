package sort

func CombSort(slice []int) {

	gap := len(slice)

	swapped := true

	for gap > 1 || swapped {

		if gap > 1 {

			gap = int(float32(gap) / 1.24733)
		}

		i := 0

		swapped = false

		for i+gap < len(slice) {

			if slice[i] > slice[i+gap] {

				slice[i], slice[i+gap] = slice[i+gap], slice[i]

				swapped = true
			}

			i++
		}
	}
}
