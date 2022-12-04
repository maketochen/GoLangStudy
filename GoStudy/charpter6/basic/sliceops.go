package main

import "fmt"

func main() {
	var s []int //Zero value for slice is nil

	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}

	//fmt.Println(s)
	s1 := make([]int, 16, 32)
	fmt.Println(s1)
	printSlice(s1)

}

func printSlice(s []int) {
	fmt.Printf("len=%d,cap=%d \n", len(s), cap(s))
}
