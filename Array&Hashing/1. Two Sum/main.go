package main

func main(){
	nums := []int{2, 11, 15, 7}
	target := 9
	twoSum(nums, target)
}

func twoSum(nums []int, target int) []int {
    seen := make(map[int]int) 

    for i, num := range nums {
        if index, found := seen[target - num]; found {
            return []int{index, i}
        }
        seen[num] = i 
    }

    return nil 
}