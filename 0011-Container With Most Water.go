package main

func maxArea(height []int) int {
    left := 0
    right := len(height) - 1
    res := 0
    for left < right {
        area := min(height[left], height[right]) * (right - left)
        res = max(area, res)
        if height[left] > height[right] {
            right--
        } else {
            left++
        }
    }
    return res
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}
