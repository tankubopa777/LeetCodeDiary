package main

import "fmt"

func main() {
	// merge([]int{1,2,3,0,0,0}, 3, []int{2,5,6}, 3)
	price := [5]int{1,2,3,4,5}
	fmt.Println(price)
	for i := 0 ; i < len(price) ; i ++ {
		fmt.Println(price[i])
	}

}

func merge(nums1 []int, m int, nums2 []int, n int)  {
	for i := 0 ; i <= m ; i++ {
		fmt.Println(nums1[i])
	}
}