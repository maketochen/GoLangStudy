package main

import "fmt"

func main() {
	strings := "123451234567231"
	rs := lengthOfNonRepeatingSubStr(strings)
	fmt.Println(rs)
	fmt.Println(lengthOfNonRepeatingSubStr("黑发肥发黑发灰会花飞"))
}

func lengthOfNonRepeatingSubStr(s string) int {

	last0ccured := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {

		if lastI, ok := last0ccured[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		last0ccured[ch] = i
	}
	return maxLength
}
