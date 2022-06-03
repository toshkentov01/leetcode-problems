package source

import (
	"fmt"
	"testing"
	"time"
)

type TestCase struct {
	Num1, Num2, Expected int
}

var testCases []TestCase = []TestCase{
	{
		Num1:     2,
		Num2:     3,
		Expected: 3,
	},
	{
		Num1:     10,
		Num2:     10,
		Expected: 1,
	},
}

func TestCountOperations(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Testing CountOperations function at %s", time.Now()), func(t *testing.T) {
			result := CountOperations(testCase.Num1, testCase.Num2)

			if result != testCase.Expected {
				t.Errorf("Test FAILED. WANTED: %v, GOT: %v", testCase.Expected, result)
			}
		})
	}
}
