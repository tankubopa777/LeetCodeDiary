package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(productExceptSelf([]int{1,2,3,4}))

}

func productExceptSelf(nums []int) []int {

	var prefix []int 
	stack := 1 
    for _, value := range nums {
		value = value * stack
		stack = value
		prefix = append(prefix, value)
	}

	sort.Slice(nums, func(i,j int) bool{
		return nums[i] >= nums[j]
	})
	
	var postfix []int
	stack2 := 1
	for _, value := range nums {
		value = value * stack2
		stack2 = value
		postfix = append(postfix, value)
	}

	sort.Slice(postfix, func(i,j int) bool{
		return postfix[i] >= postfix[j]
	})

	fmt.Println(prefix)
	fmt.Println(postfix)

	return prefix
}