package sort_colors

func Solve(i []int) []int {
	r, w, b := 0, 0, len(i) -1
	red, white, blue := 0, 1, 2

	for w <= b {
		if i[w] == white {
			w++
			continue
		}

		if i[w] == red {
			i[w], i[r] = i[r], i[w]
			w++
			r++
			continue
		}

		if i[w] == blue {
			i[w], i[b] = i[b], i[w]
			b--
			continue
		}
	}
	return i
}
