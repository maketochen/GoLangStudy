package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "我爱写bug"
	fmt.Println(s)

	for i, ch := range s {
		fmt.Println(i, ch)
	}
	for i, ch := range []byte(s) {
		fmt.Printf("(%d --  %c)", i, ch)
	}
	for i, ch := range []rune(s) {
		fmt.Printf("(%d --  %c)", i, ch)
	}
	fmt.Println()
	fmt.Println("Rune Count:", utf8.RuneCountInString(s))

	// bytes := []byte(s)
	// for len(bytes) > 0 {
	// 	ch, size := utf8.DecodeLastRune(bytes)
	// 	bytes = bytes[size:]
	// 	fmt.Printf("%c", ch)
	// }

}
