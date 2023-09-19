package main
func containsDuplicate(nums []int) bool {
    duplicateMap := make(map[int]struct{})
    for _, v := range nums {
        if _, ok := duplicateMap[v]; ok {
            return true
        }
        duplicateMap[v] = struct{}{}
    }
    return false
}
