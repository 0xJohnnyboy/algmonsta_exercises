package container_water

import (
	"testing"
)

type TestCase struct {
	Input    []int
	Expected int
}

var testCases = []TestCase{
	{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
	{[]int{1, 1}, 1},
	{[]int{1, 2, 1}, 2},
	{[]int{1, 2, 4, 3}, 4},
	{[]int{2, 3, 4, 5, 18, 17, 6}, 17},
	{[]int{1, 3, 2, 5, 25, 24, 5}, 24},
}

func TestSolve(t *testing.T) {
	for _, tc := range testCases {
		op := Solve(tc.Input)

		if op != tc.Expected {
			t.Errorf("Input: %v\nOutput: %v; expected %v \n", tc.Input, op, tc.Expected)
		}
	}
}
