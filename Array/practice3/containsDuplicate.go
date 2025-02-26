package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	fmt.Println(containsDuplicate(nums))
}

func containsDuplicate(nums []int) bool {
	numSet := make(map[int]struct{})

	for _, num := range nums {
		if _, exists := numSet[num]; exists {
			return true
		} else {
			numSet[num] = struct{}{}
		}
	}
	return false
}
