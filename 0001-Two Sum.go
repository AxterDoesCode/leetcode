package main

func twoSum(nums []int, target int) []int {
    checked := map[int]int{}
    for index, val := range nums {
        find := target - val
        if key, ok := checked[find]; !ok{
            checked[val] = index
        } else {
            return []int{key, index}
        }
    }
    return nil
}
