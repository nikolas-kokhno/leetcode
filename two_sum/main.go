package two_sum

func TwoSum(nums []int, target int) []int {
	seenNums := map[int]int{}

	for i, num := range nums {
		potentialRes := target - num
		if j, found := seenNums[potentialRes]; found {
			return []int{j, i}
		}
		seenNums[num] = i
	}

	return nil
}
