package main

//Calculate the total of everything before the number and then everything after
//2 Passes through the array to get the prefix and postfix which mutate the same slice of data "res"

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	prefix := 1
	for i, num := range nums {
		res[i] = prefix
		prefix *= num
	}

	postfix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= postfix
		postfix *= nums[i]
	}
	return res
}
