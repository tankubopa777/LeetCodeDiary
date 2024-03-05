package main

import (
	"fmt"
	"sort"
	"strings"
)



func main() {
    
	// // covert int to string
	// var i int = 42
	// var j string = strconv.Itoa(i)
	// fmt.Println(j)

	// // convert string to int
	// var k string = "42"
	// l, _ := strconv.Atoi(k)
	// fmt.Println(l)

	// // sort string
	// var m string = "Hello"
	// n := []byte(m)
	// fmt.Println(n)

	// // sort list int
	// o := []int{1, 6, 7, 4, 5}
	// sort.Ints(o)
	// p := o
	// fmt.Println(p)

	// findTheDifference("abcdefg","abcdefgf")
	// containsDuplicate([]int{5,9,8,4,4})


}

func findTheDifference(s string, t string) {
	var result string
	for _, v := range t {
		println(string(v))
		println(v)
		if strings.Contains(s, string(v)) {
			s = strings.Replace(s, string(v), "", 1)
		} else {
			result = string(v)
		}
	}
	fmt.Println(result)
}

func containsDuplicate(nums []int) bool {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] <= nums[j]
	})

    for i := 0 ; i < len(nums); i++ {
		if i < (len(nums) - 1){
		if nums[i] == nums[i+1]  {
			return true
		}}
	}

	return false
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