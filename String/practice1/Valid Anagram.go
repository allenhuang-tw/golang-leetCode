package main

import (
	"fmt"
	"sort"
)

func main() {
	anwser := isAnagram("anagram", "nagaram")
	fmt.Print(anwser)
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sArr := []rune(s)
	tArr := []rune(t)

	sort.Slice(sArr, func(i int, j int) bool { return sArr[i] < sArr[j] })
	sort.Slice(tArr, func(i int, j int) bool { return tArr[i] < tArr[j] })

	if string(sArr) == string(tArr) {
		return true
	} else {
		return false
	}
}
