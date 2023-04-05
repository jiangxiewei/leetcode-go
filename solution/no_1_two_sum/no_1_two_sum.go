package main

import "encoding/json"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := range nums {
		if t, exist := m[target-nums[i]]; exist {
			return []int{i, t}
		}
		m[nums[i]] = i
	}
	return nil
}

func main() {
	marshal, _ := json.Marshal(twoSum([]int{2, 7, 11, 15}, 9))
	println(string(marshal))
}
