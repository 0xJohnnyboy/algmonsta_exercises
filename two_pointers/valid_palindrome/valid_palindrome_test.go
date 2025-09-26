package valid_palindrome

import (
	"testing"
)

type TestCase struct {
	Input  string
	Expected bool
}

var testCases = []TestCase{
	{Input: "A man, a plan, a canal: Panama", Expected: true},
	{Input: "race a car", Expected: false},
	{Input: "  ", Expected: true},
	{Input: "cool, look ! kool, lock", Expected: false},
	{Input: "cool, look ! kool, looc", Expected: true},
}

func TestSolve(t *testing.T) {
	for _, tc := range testCases {
		op := Solve(tc.Input)

		if tc.Expected != op {
		t.Errorf(`Output: %v; expected %v`, op, tc.Expected)
			return
		}
	}
}
