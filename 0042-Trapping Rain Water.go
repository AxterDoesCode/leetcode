package main

func trap(height []int) int {
   l, r := 0, len(height) - 1 
   lMax, rMax := height[l], height[r]
   res := 0
   for l < r {
       if lMax < rMax{
           l++
           lMax = max(height[l], lMax)
           res += lMax - height[l]
       } else {
           r--
           rMax = max(height[r], rMax)
           res += rMax - height[r]
       }
   }
   return res
}

func max (a, b int) int {
    if a < b {
        return b
    }
    return a
}
