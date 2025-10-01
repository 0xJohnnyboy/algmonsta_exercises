package fruit_baskets

import (
	"testing"
)

type TestCase struct {
	Input    []int
	Expected int
}

var testCases = []TestCase{
	{[]int{1, 2, 1, 2, 2}, 5},
	{[]int{0, 1, 2, 2}, 3},
	{[]int{1, 2, 3, 2, 2}, 4},
	{[]int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}, 5},
	{[]int{1, 1, 1, 1}, 4},
	{[]int{1, 2, 3, 4, 5}, 2},
	{[]int{0, 1, 6, 6, 4, 4, 6}, 5},
	{[]int{1, 0, 1, 4, 1, 4, 1, 2, 3}, 5},
}

func TestSolve(t *testing.T) {
	for _, tc := range testCases {
		op := Solve(tc.Input)

		if op != tc.Expected {
			t.Errorf("❌ | Input: %v, Output: %v; expected %v \n", tc.Input, op, tc.Expected)
		}
	}
}
