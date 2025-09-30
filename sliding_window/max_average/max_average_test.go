package max_average

import (
	"math"
	"testing"
)

type TestCase struct {
	Nums     []int
	K        int
	Expected float64
}

var testCases = []TestCase{
	{[]int{1, 12, -5, -6, 50, 3}, 4, 12.75},
	{[]int{5}, 1, 5.0},
	{[]int{0, 1, 1, 3, 3}, 4, 2.0},
	{[]int{4, 0, 4, 3, 3}, 5, 2.8},
	{[]int{-1}, 1, -1.0},
	{[]int{1, 2, 3, 4}, 2, 3.5},
	{[]int{9, 7, 3, 5, 6, 9, 8, 5, 6, 9}, 6, 7.16667},
}

func TestSolve(t *testing.T) {
	for _, tc := range testCases {
		op := Solve(tc.Nums, tc.K)

		if math.Abs(op-tc.Expected) > 1e-5 {
			t.Errorf("❌ | Input: %v with k:%v\nOutput: %v; expected %v \n", tc.Nums, tc.K, op, tc.Expected)
		}
	}
}
