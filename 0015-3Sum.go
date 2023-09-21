package main

import (
    "sort"
)

func threeSum(nums []int) [][]int {
    res := [][]int{}
    //Sort the array
    sort.Ints(nums)
    for i, v := range nums {
        //Skip all duplicates from the left
        //i > 0 ensures that this check is done from the 2nd element onwards
        if i > 0 && v == nums[i - 1] {
            continue
        }
        l,r := i + 1, len(nums) - 1
        for l < r {
            sum := v + nums[l] + nums[r]
            if sum > 0 {
                r--
            } else if sum < 0 {
                l++
            } else {
                res = append(res, []int{v,nums[l],nums[r]})
                l++
                //Skip all the duplicates from the left
                for nums[l] == nums[l-1] && l < r {
                    l++
                }
            }
        }
    }
    return res
}
