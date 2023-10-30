package climbstairs

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	dp := make([]int, n+1)

	if len(dp) < 2 {
		return len(dp)
	}

	dp[1], dp[2] = 1, 2

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-2] + dp[i-1]
	}

	return dp[n]
}
