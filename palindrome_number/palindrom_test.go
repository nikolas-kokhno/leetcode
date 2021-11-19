package palindromenumber_test

import (
	"fmt"
	"testing"

	palindromenumber "github.com/nikolas-kokhno/leetcode/palindrome_number"
)

func TestFindDisappearedNumbers(t *testing.T) {
	for i, tt := range []struct {
		num int
		res bool
	}{
		{
			121,
			true,
		},
		{
			322,
			false,
		},
		{
			0,
			true,
		},
		{
			-232,
			false,
		},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			res := palindromenumber.IsPalindrome(tt.num)
			if tt.res != res {
				t.Errorf("want %v; got %v", tt.res, res)
			}
		})
	}
}
