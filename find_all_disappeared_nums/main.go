package findalldisappearednums

import "sort"

/*
	Given an array nums of n integers where nums[i] is in the range [1, n],
	return an array of all the integers in the range [1, n] that do not appear in nums.
*/

func FindDisappearedNumbers(nums []int) []int {
	var res []int
	var key int = 1
	sort.Ints(nums)

	numsMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		numsMap[nums[i]] = nums[i]
	}

	for i := 0; i < len(nums); i++ {
		if key != numsMap[key] {
			res = append(res, key)
		}
		key++
	}

	return res
}
