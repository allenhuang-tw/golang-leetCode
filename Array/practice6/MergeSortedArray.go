package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	m := 3
	n := 3
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {

	k := m - 1
	j := n - 1

	for i := m + n - 1; i >= 0; i-- {
		if k < 0 {
			nums1[i] = nums2[j]
			j--
			continue
		} else if j < 0 {
			nums1[i] = nums1[k]
			k--
			continue
		}
		if nums1[k] > nums2[j] {
			nums1[i] = nums1[k]
			k--
		} else {
			nums1[i] = nums2[j]
			j--
		}
	}
}
