package main

func groupAnagrams(strs []string) [][]string {
	anagramMap := make(map[[26]int][]string)
	for _, v := range strs {
		var count [26]int
		for _, c := range v {
			count[c-'a']++
		}
		anagramMap[count] = append(anagramMap[count], v)
	}
	result := make([][]string, len(anagramMap))
	i := 0
	for _, v := range anagramMap {
		result[i] = v
		i++
	}
	return result
}
