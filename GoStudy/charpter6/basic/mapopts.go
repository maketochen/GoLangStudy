package main

import "fmt"

func main() {
	strings := "123451234567231"
	rs := lengthOfNonRepeatingSubStr(strings)
	fmt.Println(rs)
	fmt.Println(lengthOfNonRepeatingSubStr("这是一个练习"))
}

func lengthOfNonRepeatingSubStr(s string) int {

	last0ccured := make(map[byte]int)
	start := 0
	maxLength := 0

	for i, ch := range []byte(s) {

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
