package main

import "fmt"

func main() {
	planets := [...]string{
		"1",
		"2",
		"3",
	}
	terraform(&planets)
	fmt.Println(planets)
}

func terraform(planets *[3]string) {
	for i := range planets {
		planets[i] = "New  " + planets[i]
	}
}
