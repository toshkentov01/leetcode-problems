package source

func searchRange(nums []int, target int) []int {

	var (
		start, end int = -1, -1
		found      bool
	)

	for index, num := range nums {
		if num == target {
			if found {
				end = index
				continue
			}

			start = index
			found = true
		}
	}

	if found {
		if end != -1 {
			return []int{start, end}
		}

		return []int{start, start}
	}

	return []int{-1, -1}
}
