package two_sum_2

func Solve(input []int, target int) ([2]int) {
	left, right := 0, len(input)-1

	for left < right {
		sum := input[left] + input[right]
		if sum == target {
			return [2]int{left + 1, right + 1}
		}

		if sum < target {
			left++
		} else {
			right--
		}
	}

	return [2]int{}
}
