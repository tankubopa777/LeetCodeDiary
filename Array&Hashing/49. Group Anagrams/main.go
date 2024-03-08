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

	// seen := make(map[int]int)

	// for _ , word := range strs {
	// 	wSlice := []rune(word)

	// 	sort.Slice(wSlice, func(i, j int)  bool {
	// 		return wSlice[i] <= wSlice[j]
	// 	})

	// 	if 

	// }

	// sort.Slice( strs, func(i, j int) bool {
	// 	return strs[i] <= strs[j]
	// })

	fmt.Println(strs)

	return [][]string {{"hello","yoyoy"},{"yoyo"}}

}

