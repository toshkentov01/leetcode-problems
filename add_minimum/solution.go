package addminimum

func addMinimum(word string) int {
	var (
		count, i, k int
	)

	const (
		pattern = "abc"
	)

	for i < len(word) {
		if k == 3 {
			k = 0
		}

		if pattern[k] != word[i] {
			count++
			k++
			continue
		}

		k++
		i++
	}

	count += extraCount(word[len(word)-1])

	return count
}

func extraCount(letter byte) int {
	if letter == 'a' {
		return 2
	}

	if letter == 'b' {
		return 1
	}

	return 0
}
