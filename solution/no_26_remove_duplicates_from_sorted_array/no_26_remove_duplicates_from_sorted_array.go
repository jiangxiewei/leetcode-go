package main

import "leetcode-go/tools/json"

func removeDuplicates(nums []int) int {
	m := make(map[int]bool)
	var appendPtr = 0
	for i := range nums {
		if _, exist := m[nums[i]]; exist {
			continue
		} else {
			nums[appendPtr] = nums[i]
			appendPtr++
			m[nums[i]] = true
		}
	}
	return appendPtr
}

func main() {
	ints := []int{1, 1, 2}
	println("k:", removeDuplicates(ints), ",arr:", json.ToString(ints))
	ints = []int{1, 2, 1, 2, 3}
	println("k:", removeDuplicates(ints), ",arr:", json.ToString(ints))

}
