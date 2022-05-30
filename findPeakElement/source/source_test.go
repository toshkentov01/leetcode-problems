package source

import (
	"fmt"
	"testing"
	"time"
)

type TestCase struct {
	Nums     []int
	Expected int
}

func TestFindPeakElement(t *testing.T) {
	testCases := []TestCase{
		{
			Nums:     []int{1, 2, 3, 2, 1},
			Expected: 2,
		},
		{
			Nums:     []int{3, 2, 1, 0},
			Expected: 0,
		},
		{
			Nums:     []int{1, 2, 3, 4, 5},
			Expected: 4,
		},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Testing FindPeakElement method at: %s", time.Now()), func(t *testing.T) {
			result := FindPeakElement(testCase.Nums)
			
			if result != testCase.Expected {
				t.Errorf("Error while testing FindPeakElement method. Got %v. Expected %v", result, testCase.Expected)
			}
		})
	}
}
