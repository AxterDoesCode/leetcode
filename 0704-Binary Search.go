package main

func search(nums []int, target int) int {
    l := 0
    h := len(nums) -1

    for l <= h {
        mid := (l + h)/2
        if nums[mid] == target {
            return mid
        } else if nums[mid] < target{
            l = mid + 1
        } else {
            h = mid - 1
        }
    }
    return -1
}
