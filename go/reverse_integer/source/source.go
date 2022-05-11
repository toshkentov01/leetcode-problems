package source

func Reverse(x int) int {

	border := 1 << 31
	reversed := 0
	positive := true

	if x < 0 {
		positive = false
		x = x * -1
	}

	for x > 0 {
		mod := x % 10
		reversed = reversed*10 + mod

		if reversed > border || reversed < -border {
			return 0
		}
		x /= 10
	}

	if !positive {
		return reversed * -1
	}

	return reversed
}