package contains_duplicate

func Solve(nums []int, k int) bool {
	if k == 0 {
		return false
	}

	seenSet := make(map[int]struct{})
	for i, num := range nums {
		if _, ok := seenSet[num]; ok {
			return true
		}

		seenSet[num] = struct{}{} // using structs for set because empty struct occupy no memory

		if i >= k {
			delete(seenSet, nums[i-k])
		}
	}
	return false
}
