package main

import (
	"fmt"
	"math/rand"
	"time"
)

func msgGen(name string, done chan struct{}) chan string {
	c := make(chan string)
	go func() {
		i := 0
		for {
			select {
			case <-time.After(time.Duration(rand.Intn(2000)) * time.Millisecond):
				c <- fmt.Sprintf("service %s: message: %d", name, i)
			case <-done:
				fmt.Println("cleaning up")
				time.Sleep(2 * time.Second)
				fmt.Println("cleaning done")
				done <- struct{}{}
				return
			}
			i++
		}
	}()
	return c
}

func nonBlockingWait(c chan string) (string, bool) {
	select {
	case m := <-c:
		return m, true
	default:
		return "false", false
	}
}

func fanIn(chs ...chan string) chan string {
	c := make(chan string)
	for _, ch := range chs {
		go func(in chan string) {
			for {
				c <- <-in
			}
		}(ch)
	}
	return c
}

//多路复用
func fanInSelect(c1, c2 chan string) chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case m := <-c1:
				c <- m
			case m := <-c2:
				c <- m
			}
		}
	}()
	return c
}

func timeoutWait(c chan string, timeout time.Duration) (string, bool) {
	select {
	case m := <-c:
		return m, true
	case <-time.After(timeout):
		return "false", false
	}
}

func main() {
	//并发模式
	//m1 := msgGen("service1")
	//m2 := msgGen("service2")
	//m := fanIn(m1, m2)
	//m := fanInSelect(m1, m2)
	done := make(chan struct{})
	m1 := msgGen("service1", done)
	for i := 0; i < 5; i++ {
		//fmt.Println("m1", <-m1)
		//fmt.Println("m2", <-m2)
		//fmt.Println(<-m)
		//if m, ok := nonBlockingWait(m2); ok {
		//	fmt.Println(m)
		//} else {
		//	fmt.Println("no message from service2")
		//}
		if m, ok := timeoutWait(m1, time.Second); ok {
			fmt.Println(m)
		} else {
			fmt.Println("timeout")
		}
	}
	done <- struct{}{}
	<-done
}
