package source

import (
	"fmt"
	"testing"
	"time"
)

type testCase struct {
	prices   []int
	expected int
}

func TestMaxProfit(t *testing.T) {
	testCases := []testCase{
		{
			prices:   []int{7, 1, 5, 6, 4},
			expected: 5,
		},
		{
			prices:   []int{1},
			expected: 0,
		},
	}

	t.Run(fmt.Sprintf("Testing method MaxProfit at: %s", time.Now()), func(t *testing.T) {
		for _, testCase := range testCases {
			if result := MaxProfit(testCase.prices); result != testCase.expected {
				t.Errorf("FAIL: WANT: %v, GOT: %v", testCase.expected, result)
			}
		}
	})
}
