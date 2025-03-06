package main

import "fmt"

func main() {
	str := "pwwkew"

	ans := lengthOfLongestSubstring(str)

	fmt.Print(ans)
}

func lengthOfLongestSubstring(s string) int {
	charIndexMap := make(map[rune]int)
	maxLength := 0
	left := 0

	for right, char := range s {
		if prevIndex, found := charIndexMap[char]; found && prevIndex >= left {
			left = prevIndex + 1
		}

		charIndexMap[char] = right
		currentMaxLength := right - left + 1

		if currentMaxLength > maxLength {
			maxLength = currentMaxLength
		}

	}

	return maxLength
}
