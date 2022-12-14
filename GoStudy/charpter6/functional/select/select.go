package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var c1, c2 = generator(), generator()
	//w := createWork(0)
	var worker = createWork(0)
	var values []int
	tm := time.After(10 * time.Second) //10s 后退出
	tick := time.Tick(time.Second)     //计时
	n := 0
	//hasValue := false
	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		select {
		case n = <-c1:
			//fmt.Println("Received: from c1", w)
			//w <- n
			//hasValue = true
			values = append(values, n)
		case n = <-c2:
			//fmt.Println("Received: from c2", n)
			//w <- n
			//hasValue = true
			values = append(values, n)
		case activeWorker <- activeValue:
			//hasValue = false
			values = values[1:]
		case <-time.After(800 * time.Millisecond):
			fmt.Println("timeout")
		case <-tick:
			fmt.Println("queue len=", len(values))
		case <-tm:
			fmt.Println("10s,bye")
			return
		}
	}
}

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(
				time.Duration(rand.Intn(1500)) *
					time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}
func work(id int, c chan int) {
	for n := range c {
		time.Sleep(time.Second * 1)
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}
func createWork(id int) chan<- int {
	c := make(chan int)
	go work(id, c)
	return c
}
