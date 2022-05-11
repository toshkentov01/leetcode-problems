package source

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

type TestCase struct {
	Nums     []int
	Expected int
}

func TestMajorityElement(t *testing.T) {
	testCases := []TestCase{
		{
			Nums:     []int{3, 2, 3},
			Expected: 3,
		},
		{
			Nums:     []int{1, 2, 2, 3, 2},
			Expected: 2,
		},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Testing MajorityElement method at: %s", time.Now()), func(t *testing.T) {
			result := MajorityElement(testCase.Nums)

			if difference := cmp.Diff(result, testCase.Expected); difference != "" {
				t.Error(difference)
			}
		})
	}
}
