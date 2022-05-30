package source

import (
	"fmt"
	"testing"
	"time"
)

type TestCase struct {
	String   string
	Letter   byte
	Expected int
}

func TestPercentageLetter(t *testing.T) {
	testCases := []TestCase{
		{
			String:   "foobar",
			Letter:   byte('o'),
			Expected: 33,
		},
		{
			String:   "jjjj",
			Letter:   byte('a'),
			Expected: 0,
		},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Testing PercentageLetter method at %s", time.Now()), func(t *testing.T) {
			result := PercentageLetter(testCase.String, testCase.Letter)

			if result != testCase.Expected {
				t.Errorf("Error while testing PercentageLetter method. Got %v, wanted %v", result, testCase.Expected)
			}
		})
	}
}
