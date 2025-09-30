package sort_colors

import (
	"testing"
)

type TestCase struct {
	Input    []int
	Expected []int
}

var testCases = []TestCase{
	{[]int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
	{[]int{2, 0, 1}, []int{0, 1, 2}},
	{[]int{0}, []int{0}},
	{[]int{1}, []int{1}},
	{[]int{2}, []int{2}},
	{[]int{1, 2, 0}, []int{0, 1, 2}},
	{[]int{2, 2, 2, 0, 0, 0, 1, 1}, []int{0, 0, 0, 1, 1, 2, 2, 2}},
	{[]int{1, 0, 1, 0, 1, 0}, []int{0, 0, 0, 1, 1, 1}},
}

func TestSolve(t *testing.T) {
	for _, tc := range testCases {
		op := Solve(tc.Input)

		if ok := areSlicesEqual(tc.Expected, op); !ok {
			t.Errorf("Input was: %v; Output: %v; expected %v \n",tc.Input, op, tc.Expected)
		}
	}
}

func areSlicesEqual(s1 []int, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}
