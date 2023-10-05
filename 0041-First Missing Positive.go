package main

import "fmt"

func firstMissingPositive(nums []int) int {
	n := len(nums)
	//While loop to ensure that a number is inserted into its sorted position
	for i := 0; i < n; {
		if nums[i] > n || nums[i] <= 0 || nums[i] == nums[nums[i]-1] {
			i++
			continue
		}
		nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
	}
	fmt.Println(nums)
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}
