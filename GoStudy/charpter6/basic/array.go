package main

import "fmt"

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var grid [4][5]int

	fmt.Println(arr3[2:])
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)
	s1 := arr3[2:6]
	s2 := s1[3:6]
	s3 := append(s2, 19)
	s4 := s3[:]
	fmt.Println(s3, s4)
	fmt.Println(s1, s2)
}
