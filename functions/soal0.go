package functions

import (
	"math"
	"strconv"
	"strings"
)

func strRev(str string) string {
	newStr := ""
	for i := len(str); i != 0; i-- {
		newStr += string(str[i-1])
	}
	return newStr
}

func IsPalindrome(val int) bool {
	str1 := strconv.Itoa(val)
	str2 := strRev(str1)
	if strings.Compare(str1, str2) == 0 {
		return true
	} else {
		return false
	}
}

func GetPalindromes(from, to int) []int {
	res := make([]int, 0)
	if from < 1 || to > int(math.Pow(10, 9)) {
		return nil
	}
	for i := from; i <= to; i++ {
		if IsPalindrome(i) {
			res = append(res, i)
		}
	}
	return res
}
