package main

import (
	"fmt"
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

func lengthOfLongestSubstring3(s string) int {
    //This bool key is important because its used as the condition for the nested for loop
    charSet := make(map[byte]bool)
    l := 0 //Left pointer
    res := 0
    
    //r = right pointer
    for r := range s {
        for charSet[s[r]] {
            delete(charSet,s[l])
            l++
        }
        fmt.Println(charSet)
        charSet[s[r]] = true
        res = max(res, r-l+1)
    }
    return res
}

func max(a,b int) int {
    if a > b{
        return a
    }
    return b
}

