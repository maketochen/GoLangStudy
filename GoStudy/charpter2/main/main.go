package main

// func main() {

// 	arr := [...]int{1, 2, 4, 3, 5, 8}
// 	fmt.Println("排序后", arr)
// 	bubblesort(&arr)
// 	fmt.Println("hello world")
// 	nm()
// }

// func bubblesort(arr *[6]int) {
// 	var temp = 0
// 	for i, it := 0, len(*arr); i < it-1; i++ {
// 		for j := 0; j < it-1-i; j++ {
// 			if arr[j] > arr[j+1] {
// 				temp = arr[j+1]
// 				arr[j+1] = arr[j]
// 				arr[j] = temp
// 			}
// 		}
// 	}
// 	fmt.Println("排序后", *arr)
// }

// //go tips
// func nm() {
// 	a := [3]int{0, 1, 2}
// 	s := a[1:2]
// 	s[0] = 11
// 	s = append(s, 12)
// 	s = append(s, 13)
// 	s[0] = 21
// 	fmt.Println(a)
// 	fmt.Println(s)
// }

type People interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {

}

func main() {

	//var s *Student
	//if s == nil {
	//	fmt.Println("s is nil")
	//} else {
	//	fmt.Println("s is not nil")
	//}
	//var p People = s
	//if p == nil {
	//	fmt.Println("p is nil")
	//} else {
	//	fmt.Println("p is not nil")
	//}
	var a *int
	*a += 1
	// DoOperation(1, increase)
	DoOperation(1, decrease)
}
func increase(a, b int) int {
	return a + b
}

func DoOperation(y int, f func(int, int)) {
	f(y, 1)
}

func decrease(a, b int) {
	println("decrease result is:", a-b)
}
