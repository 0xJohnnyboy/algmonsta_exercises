package two_pointers

import (
	"testing"
)

type TestCase struct {
	Input  []int
	Target int
	Output [2]int
}

var testCases = []TestCase{
	{Input: []int{2, 7, 11, 15}, Target: 9, Output: [2]int{1, 2}},
	{Input: []int{2, 3, 4}, Target: 6, Output: [2]int{1, 3}},
	{Input: []int{3, 8, 15, 35, 43, 46, 48, 52}, Target: 87, Output: [2]int{4, 8}},
}

func TestTwoSum(t *testing.T) {
	for _, tc := range testCases {
		op := TwoSum2(tc.Input, tc.Target)

		if tc.Output[0] != op[0] {
			t.Errorf(`Output[0] %d doesn't match expected %d`, op[0], tc.Output[0])
			return
		}
		if tc.Output[1] != op[1] {
			t.Errorf(`Output[1] %d doesn't match expected %d`, op[1], tc.Output[1])
			return
		}
	}
}
