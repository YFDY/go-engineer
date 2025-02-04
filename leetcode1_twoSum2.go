package main

import "fmt"

func twoSum2(nums []int, target int) []int {
	hashTable := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		p, ok := hashTable[target-nums[i]]
		if ok {
			return []int{i, p}
		}
		hashTable[nums[i]] = i
	}
	return nil
}

func main() {
	a := []int{2, 7, 11, 15}
	target := 9
	result := twoSum2(a, target)
	fmt.Println(result)
}

/*
p, ok := hashTable[target-nums[i]]  // 如果键不存在，ok 的值为 false， p的值为该类型的零值
时间复杂度O(n),对于每一个元素 x，我们可以 O(1)O(1)O(1) 地寻找 target - x。
空间复杂度O(n),主要为哈希表的开销
*/
