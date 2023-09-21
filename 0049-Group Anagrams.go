package main

/*Fairly simple problem, LC Medium, get the count of each letter and use the corresponding index
as part of a count array. Use this array as the key in a map and then append the current string to the key
of the map

Then create a 2d array for the results and append the anagramMap contents into it
Note that the length of the result 2d array will be the same as the length of the anagramMap aka the number of
keys in the map*/

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
