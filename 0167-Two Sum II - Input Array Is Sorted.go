package main

func twoSum2(numbers []int, target int) []int {
   //You solve this like a binary search 
   p1 := 0
   p2 := len(numbers) - 1
   for p1 < p2 {
       sum := numbers[p1] + numbers[p2]
       if sum == target {
           return []int{p1 + 1, p2 + 1}
       } else if sum > target {
           p2--
       } else {
           p1++
       }
   }
   return []int{0,0}
}
