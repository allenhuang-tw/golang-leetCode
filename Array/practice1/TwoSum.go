package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		sum := target - num
		if index, found := numMap[sum]; found {
			return []int{index, i}
		}

		numMap[num] = i
	}
	return nil
}
