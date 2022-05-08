package source

import (
	"testing"
)

func TestSearch(t *testing.T) {
	type TestCase struct {
		nums     []int
		target   int
		expected int
	}

	testCases := []TestCase{
		{
			nums:     []int{1, 2, 3},
			target:   2,
			expected: 1,
		},
		{
			nums:     []int{1, 11, 2, 4},
			target:   1,
			expected: 0,
		},
		{
			nums:     []int{1, 2, 3},
			target:   4,
			expected: -1,
		},
	}

	t.Run("Testing the Search method", func(t *testing.T) {
		for _, testCase := range testCases {
			result := Search(testCase.nums, testCase.target)

			if result != testCase.expected {
				t.Errorf("The result is %d: . But wanted value was: %d", result, testCase.expected)
			}
		}
	})
}
