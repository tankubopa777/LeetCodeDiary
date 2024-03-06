package main

import (
	"fmt"
	"sort"
)

func main(){
	fmt.Println(isAnagram("anagram","nagaram"))
}

func isAnagram(s string, t string) bool {
	sSlice := []rune(s)
	tSlice := []rune(t)
	sort.Slice(sSlice, func(i, j int) bool {
		return sSlice[i] < sSlice[j]
	})
	sort.Slice(tSlice, func(i, j int) bool {
		return tSlice[i] < tSlice[j]
	})

	return string(sSlice) == string(tSlice) 
}

func twoSum(nums []int, target int) []int {
    
}