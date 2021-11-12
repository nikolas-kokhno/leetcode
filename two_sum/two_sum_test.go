package two_sum_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/nikolas-kokhno/leetcode/two_sum"
)

func TestTwoSum(t *testing.T) {
	for i, tt := range []struct {
		nums   []int
		target int
		result []int
	}{
		{
			[]int{1, 2, 3, 4},
			5,
			[]int{1, 2},
		},
		{
			[]int{2, 7, 11, 15},
			9,
			[]int{0, 1},
		},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			res := two_sum.TwoSum(tt.nums, tt.target)
			if !reflect.DeepEqual(tt.result, res) {
				t.Errorf("want %v; got %v", tt.result, res)
			}
		})
	}
}
