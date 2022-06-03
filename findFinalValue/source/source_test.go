package source

import (
	"fmt"
	"testing"
	"time"
)

type TestCase struct {
	Nums     []int
	Original int
	Expected int
}

var testCases []TestCase = []TestCase{
	{
		Nums:     []int{5, 3, 6, 1, 12},
		Original: 3,
		Expected: 24,
	},
	{
		Nums:     []int{5, 3, 4, 7, 2, 9},
		Original: 2,
		Expected: 8,
	},
}

func TestFindFinalValue(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Testing FindFinalValue function at %s", time.Now()), func(t *testing.T) {
			result := FindFinalValue(testCase.Nums, testCase.Original)

			if result != testCase.Expected {
				t.Errorf("Test FAILED. WANTED: %v, GOT: %v", testCase.Expected, result)
			}
		})
	}
}
