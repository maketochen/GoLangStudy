package main

import "fmt"

func main() {
	var a = [5]string{"I", "am", "steup", "and", "work"}
	for b := 0; b < len(a); b++ {
		if b == 2 {
			a[b] = "smart"
			continue
		}
		if b == 4 {
			a[b] = "strong"
		}
	}
	fmt.Println(a)
}
