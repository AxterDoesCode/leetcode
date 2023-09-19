package main

func topKFrequent(nums []int, k int) []int {
    countMap := make(map[int]int)
		countSlice := make([][]int, len(nums) + 1)

		//Get the count of each number and store it in a map

		for _, v := range nums {
			if _, ok := countMap[v]; ok {
				countMap[v]++
			}else{
				countMap[v] = 1
			}
		}

		// iterate through the map and 

		for num, count := range countMap{
			countSlice[count] = append(countSlice[count], num)
		}

		res := []int{}

		for i := len(countSlice) - 1; i > 0; i-- {
			res = append(res, countSlice[i]...)
			if len(res) == k{
				return res
			}
		}
		return nil
}

//Algorithim can be improved significantly with quickSelect
