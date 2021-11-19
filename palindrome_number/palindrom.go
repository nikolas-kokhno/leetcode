package palindromenumber

/*
	Given an integer x, return true if x is palindrome integer.

	An integer is a palindrome when it reads the same backward as forward.
	For example, 121 is palindrome while 123 is not.
*/

func IsPalindrome(x int) bool {
	var value int = x

	if x < 0 {
		return false
	}

	if x == 0 || x < 10 {
		return true
	}

	newInt := 0
	for x > 0 {
		newInt = (newInt * 10) + x%10
		x /= 10
	}

	return newInt == value
}
