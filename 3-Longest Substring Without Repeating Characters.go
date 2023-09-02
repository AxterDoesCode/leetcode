package main

import (
	"math"
	"strings"
)

//This solution uses sliding window techinique
func lengthOfLongestSubstring(s string) int {
    longest := 0
    start := 0
    charIndexMap := make(map[uint8]int)
    n := len(s)
    for end := 0; end < n; end++ {
        duplicateIndex, isDuplicate := charIndexMap[s[end]]
        //If it is a duplicate calculate
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
//This one is crazy
//It is sliding window but it uses the strings std lib to do the checking
func lengthOfLongestSubstring2(s string) int {
    var highest, distance int
    prevPos := -1
    for i, v := range s {
        c := string(v)
        pos := strings.LastIndex(s[:i], c)
        if pos > prevPos {
            prevPos = pos
        }
        distance = i - prevPos
        if distance > highest {
            highest = distance
        }
    }
    return highest
}
