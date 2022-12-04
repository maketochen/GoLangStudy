package main

import "fmt"

type aint int

func main() {
	var room = "lake"

	switch {
	case room == "tuks":
		fmt.Println("this tuks")
		fallthrough
	case room == "lake":
		fmt.Println("this is lake")
		fallthrough
	case room == "fksa":
		fmt.Println("this is fksa")
	}
	var ia aint = 3
	fmt.Println(ia.sum())
}

func (i aint) sum() aint {
	i = i + 1
	return i
}
