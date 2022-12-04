package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

const filename = "cc.txt"

func main() {
	fmt.Println("123")
	variableZeroValue()
	variableInitialValue()
	euler()
	enums()
}

func euler() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
	cc := cmplx.Pow(math.E, 1i*math.Pi) + 1
	fmt.Printf("%.3f\n", cc)
}

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func enums() {
	const (
		cpp  = iota
		java = 1
	)
	fmt.Println(cpp, java)
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
	)
	fmt.Println(b, kb, mb, gb, tb)
}
