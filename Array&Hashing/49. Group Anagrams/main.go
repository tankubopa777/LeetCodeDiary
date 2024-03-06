package main

import (
	"fmt"
)

func main(){
	fmt.Println(groupAnagrams([]string{"eat","tea","tan","ate","nat","bat"}))
}

func groupAnagrams(strs []string) [][]string {
	for _, word := range strs {
		fmt.Println(word)
	}
	return [][]string{{"tansan","sanan"},{"hahaha","eieie","yoyoyo"}}

}

