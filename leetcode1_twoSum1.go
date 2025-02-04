package main

import "fmt"

func twoSum1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func main() {
	a := []int{3, 3}
	result := twoSum1(a, 6)
	fmt.Println(result)
}

/*
时间复杂度O(n2)，其中n是数组中的元素数量
空间复杂度O(1)
*/
