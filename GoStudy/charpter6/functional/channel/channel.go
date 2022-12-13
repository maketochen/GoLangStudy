package main

import (
	"fmt"
	"time"
)

func main() {
	//ChanDemo()
	bufferedChannel()
}

func createWork(id int) chan<- int {
	for {
		c := make(chan int)

		go func() {
			for {
				fmt.Printf("Worker %d received %c\n", id, <-c)
			}
		}()
		return c
	}
}

func ChanDemo() {
	//c := make(chan int)
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWork(i)
	}

	//c <- 1
	//c <- 2
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
}
