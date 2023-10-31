package climbstairs

func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}

	var result int

	a, b := 1, 1

	for i := 2; i <= n; i++ {
		result = a + b
		a, b = b, result
	}

	return result
}

