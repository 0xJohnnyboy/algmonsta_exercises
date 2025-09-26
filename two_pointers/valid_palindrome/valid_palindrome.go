package valid_palindrome

import (
	"strings"
	"unicode"
)

func Solve(input string) bool {
	runes := []rune(strings.ToLower(input))
	left, right := 0, len(runes)-1

	if len(runes) == 0 {
		return false
	}
	if len(runes) == 1 {
		return true
	}

	for left < right {
		if !IsLetterOrNumber(runes[left]) {
			left++
			continue
		} else if !IsLetterOrNumber(runes[right]) {
			right--
			continue
		} else if runes[left] != runes[right] {
			return false
		}

		left++
		right--
	}
	return true
}

func IsLetterOrNumber(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}
