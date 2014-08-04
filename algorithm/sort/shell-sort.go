package sort

func ShellSort(slice []int) {

	if len(slice) < 1 {

		return
	}

	chunk := (3 * len(slice)) + 1

	for chunk >= 1 {

		for i := chunk; i < len(slice); i++ {

			for j := i; j >= chunk && slice[j] < slice[j-chunk]; j = j - chunk {

				slice[j], slice[j-chunk] = slice[j-chunk], slice[j]
			}
		}

		chunk = chunk / 3
	}
}
