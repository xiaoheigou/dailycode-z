package problem0005

import (
	"fmt"
	"math"
)

func longestPalindrome(s string) int {

	length := len(s)
	// fmt.Println(length)
	var ss [100][100]int
	for len := 1; len <= length; len++ {
		for i := 0; i <= length-len; i++ {
			j := i + len - 1
			if i == j {
				ss[i][j] = 1
				continue
			}
			if s[i] == s[j] {
				ss[i][j] = ss[i+1][j-1] + 2
			} else {
				ss[i][j] = int(math.Max(float64(ss[i+1][j]), float64(ss[i][j-1])))
			}
		}
	}
	fmt.Println(ss[0][length-1])
	return ss[0][length-1]
}
