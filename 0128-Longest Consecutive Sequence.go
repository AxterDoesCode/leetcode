package main

//Solving this problem is fairly simple, you look at each element in a set/map
//If it it contains a number before it continue, if it doesnt it means its the start
//Of a sequence, you then iterate through that sequence and testing if the next
//Number exists, if it does the the while loop will continue incrementing a counter
//Counter is then compared and returned

func longestConsecutive(nums []int) int {
	numMap := make(map[int]bool)
	longest := 0
	//Initialise the nummap (treating it like a set)
	for _, v := range nums {
		numMap[v] = true
	}

	for i := range numMap {
		if _, ok := numMap[i-1]; !ok {
			//Start of a sequence detected
			currentLength := 1
			for numMap[i+currentLength] {
				currentLength++
			}
			longest = max(longest, currentLength)
		}
	}
	return longest
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
