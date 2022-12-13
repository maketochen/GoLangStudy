package main

import (
	"fmt"
	"queue"
)

func main() {
	q := queue.Queue{1}
	q.Push(3)
	q.Push(2)
	fmt.Println(q.Pop())
	fmt.Println(q)

}
