package main

import "math"

func lengthOfLongestSubstring(s string) int {
    longest := 0
    start := 0
    charIndexMap := make(map[uint8]int)
    n := len(s)
    for end := 0; end < n; end++ {
        duplicateIndex, isDuplicate := charIndexMap[s[end]]
        if isDuplicate {
            longest = int(math.Max(float64(longest), float64(end - start)))
            for i := start; i <= duplicateIndex; i++ {
                delete(charIndexMap, s[i])
            }
            start = duplicateIndex + 1
        }
        charIndexMap[s[end]] = end
    }
    longest = int(math.Max(float64(longest), float64(n-start)))
    return longest

}
