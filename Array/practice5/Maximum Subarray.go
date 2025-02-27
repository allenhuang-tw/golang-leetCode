package main

import "fmt"

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	maxFar := nums[0]

	for i := 1; i < len(nums); i++ {
		maxSum = max(maxSum+nums[i], nums[i])
		maxFar = max(maxSum, maxFar)
	}

	return maxFar
}
