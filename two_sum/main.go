package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

func twoSum(nums []int, target int) []int {
	seenNums := map[int]int{}

	for i, num := range nums {
		potentialRes := target - num
		if j, found := seenNums[potentialRes]; found {
			return []int{i, j}
		}
		seenNums[num] = i
	}

	return nil
}
