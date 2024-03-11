package main

import "fmt"

func main() {

	fmt.Println(productExceptSelf([]int{1,2,3,4}))

}

func productExceptSelf(nums []int) []int {

	var prefix []int 
	stack := 1 
    for _, value := range nums {
		stack = stack * value
		value = value * stack
		prefix = append(prefix, value)
	}

	return prefix
}