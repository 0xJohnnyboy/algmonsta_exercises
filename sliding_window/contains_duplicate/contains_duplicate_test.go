package contains_duplicate

import (
	"testing"
)

type TestCase struct {
	Nums     []int
	K        int
	Expected bool
}

var testCases = []TestCase{
	{[]int{1, 2, 3, 1}, 3, true},
	{[]int{1, 0, 1, 1}, 1, true},
	{[]int{1, 2, 3, 1, 2, 3}, 2, false},
	{[]int{99, 99}, 2, true},
	{[]int{1, 2, 1}, 0, false},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 9}, 3, true},
	{[]int{4, 1, 2, 3, 4}, 3, false},
	{[]int{1}, 1, false},
}

func TestSolve(t *testing.T) {
	for _, tc := range testCases {
		op := Solve(tc.Nums, tc.K)

		if op != tc.Expected {
			t.Errorf("Input: %v with k:%v\nOutput: %v; expected %v \n", tc.Nums, tc.K, op, tc.Expected)
		}
	}
}
