package main

import (
	"sort"
)



func main() {
	containsDuplicate([]int{5,9,8,4,4})

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