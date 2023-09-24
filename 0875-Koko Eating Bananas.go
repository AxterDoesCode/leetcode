package main

func minEatingSpeed(piles []int, h int) int {
   res := 1
   l := 1
   r := 10000000000
   //binary search between two values
   for l <= r {
       midK := (l + r)/2
       if canEat(piles, h, midK) {
           res = midK
           r = midK - 1
       } else {
           l = midK + 1
       }
   }
   return res
}

func canEat (piles []int, limit, k int) bool {
    totalHours := 0
    for _, v := range piles {
        //This is a thin because you need to add K ammount - 1 to get the number of hours since.
        totalHours += (v + k -1) / k
        if totalHours > limit {
            return false
        }
    }
    return true
}
