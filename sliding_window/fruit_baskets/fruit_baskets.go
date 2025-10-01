package fruit_baskets

func Solve(fruits []int) int {
	basket := make(map[int]int)
	left := 0
	maxFruits := 0

	for right := range fruits {
		basket[fruits[right]]++

		for len(basket) > 2 {
			basket[fruits[left]]--
			if basket[fruits[left]] == 0 {
				delete(basket, fruits[left])
			}
			left++
		}

		maxFruits = max(maxFruits, right-left+1)

	}
	return maxFruits
}
