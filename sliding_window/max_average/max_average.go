package max_average

func Solve(nums []int, k int) float64 {
	sum := 0
	for i := range k {
		sum += nums[i]
	}

	maxSum := sum

	for i := k; i < len(nums); i++ {

		sum = sum - nums[i-k] + nums[i]

		if sum > maxSum {
			maxSum = sum
		}

	}

	return float64(maxSum) / float64(k)
}
