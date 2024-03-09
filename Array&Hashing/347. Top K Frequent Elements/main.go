package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello golang")
}

func topKFrequent(nums []int, k int) []int {
    seen := make(map[int]int)
    for _, value := range nums {
        seen[value]++
    }

    // Create a slice of structs from the map
    type frequency struct {
        num   int
        count int
    }
    var freqs []frequency
    for num, count := range seen {
        freqs = append(freqs, frequency{num, count})
    }

    // Sort the slice based on frequency count
    sort.Slice(freqs, func(i, j int) bool {
        return freqs[i].count > freqs[j].count
    })

    // Extract the top k frequent elements
    var topK []int
    for i := 0; i < k; i++ {
        topK = append(topK, freqs[i].num)
    }
    return topK
}

func topKFrequent2(nums []int, k int) []int {
	// Creates a map where keys are int and values are int
	seen := make(map[int]int)
	for _ , value := range nums {
		seen[value]++
	}

	// Create struct
	type frequency struct {
		num int
		count int
	}

	var listOfNum []frequency
	for num, count := range seen {
		listOfNum = append(listOfNum, frequency{num, count})
	}

	sort.Slice(listOfNum , func(i , j int) bool{
		return listOfNum[i].count >= listOfNum[j].count
	})

	var listReturn []int
	for i := 0 ; i < k; i++ {
		listReturn = append(listReturn, listOfNum[i].num)
	}

	return listReturn
}



