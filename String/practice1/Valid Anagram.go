package main

import (
	"fmt"
)

func main() {
	anwser := isAnagram("anagram", "nagaram")
	fmt.Print(anwser)
}

func isAnagram(s string, t string) bool {
	//slice string and compare diffrent
	// if len(s) != len(t) {
	// 	return false
	// }

	// sArr := []rune(s)
	// tArr := []rune(t)

	// sort.Slice(sArr, func(i int, j int) bool { return sArr[i] < sArr[j] })
	// sort.Slice(tArr, func(i int, j int) bool { return tArr[i] < tArr[j] })

	// if string(sArr) == string(tArr) {
	// 	return true
	// } else {
	// 	return false
	// }

	//set
	// 如果長度不相同，直接返回 false
	if len(s) != len(t) {
		return false
	}

	// 用 map 來統計每個字母的出現次數
	counts := make(map[rune]int)

	// 遍歷第一個字串，統計每個字符的數量
	for _, char := range s {
		counts[char]++
	}

	// 遍歷第二個字串，減少每個字符的數量
	for _, char := range t {
		counts[char]--
		// 如果某個字符的數量小於 0，則不是字謎
		if counts[char] < 0 {
			return false
		}
	}

	// 如果所有字符的數量都為 0，則是字謎
	return true

}
