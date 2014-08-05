package sort

func GnomeSort(slice []int) {

	i := 1

	for i < len(slice) {

		if slice[i] >= slice[i-1] {

			i++

		} else {

			slice[i], slice[i-1] = slice[i-1], slice[i]

			if i > 1 {

				i--
			}
		}
	}
}
