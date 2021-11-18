package findalldisappearednums_test

import (
	"fmt"
	"reflect"
	"testing"

	findalldisappearednums "github.com/nikolas-kokhno/leetcode/find_all_disappeared_nums"
)

func TestFindDisappearedNumbers(t *testing.T) {
	for i, tt := range []struct {
		nums   []int
		result []int
	}{
		{
			[]int{4, 3, 2, 7, 8, 2, 3, 1},
			[]int{5, 6},
		},
		{
			[]int{1, 1},
			[]int{2},
		},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			res := findalldisappearednums.FindDisappearedNumbers(tt.nums)
			if !reflect.DeepEqual(tt.result, res) {
				t.Errorf("want %v; got %v", tt.result, res)
			}
		})
	}
}
