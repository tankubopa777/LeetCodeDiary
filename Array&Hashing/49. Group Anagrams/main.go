package main

import (
	"fmt"
	"sort"
)

func main(){
	fmt.Println(groupAnagrams([]string{"eat","tea","tan","ate","nat","bat"}))
	fmt.Println("eat" == "tea")
	
	x := []rune("eat")
	y := []rune("tea")

	sort.Slice(x, func(i, j int) bool {
		return x[i] <= x[j]
	})

	sort.Slice(y, func(q, c int) bool {
		return y[q] <= y[c]
	})

	fmt.Println(string(x) == string(y))
}

func groupAnagrams(strs []string) [][]string {
    anagramMap := make(map[string][]string) // Map to store anagrams

    for _, word := range strs {
        sortedWord := string(sortRunes(word)) // Sort runes and convert to string
        anagramMap[sortedWord] = append(anagramMap[sortedWord], word)  
    }

    result := [][]string{} // Slice to store the groups of anagrams
    for _, group := range anagramMap {
        fmt.Println(result, group)
        result = append(result, group)
    }

    return result
}

// Helper function to sort the runes of a word
func sortRunes(str string) []rune {
    runes := []rune(str)
    sort.Slice(runes, func(i, j int) bool {
        return runes[i] < runes[j]
    })
    return runes
}
