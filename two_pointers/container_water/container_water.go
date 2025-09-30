package container_water

func Solve(i []int) int {
	left, right := 0, len(i)-1
	maxArea := 0

	for left <= right {
		width := right - left
		area := getArea(width, i[left], i[right])
		if area > maxArea {
			maxArea = area
		}

		if i[left] < i[right] {
			left++
			continue
		}
		right--

	}
	return maxArea
}

func getArea(width int, height1 int, height2 int) int {
	if height1 > height2 {
		return width * height2
	}

	return width * height1
}
