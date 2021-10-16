package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var once sync.Once

func ch1(ch chan int) {
	defer close(ch)
	for i := 0; i < 100; i++ {
		ch <- i
	}
}

func ch2(ch1, ch2 chan int) {
	defer once.Do(func() { close(ch2) })
	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x * x
	}
	// for v := range ch1 {
	// 	ch2 <- v * v
	// }
}

func main() {
	c1, c2 := make(chan int, 50), make(chan int, 100)
	go ch1(c1)
	go ch2(c1, c2)
	go ch2(c1, c2)
	for v := range c2 {
		fmt.Println(v)
	}
	ch := make(chan int, 3)
	for {
		i := rand.Intn(10)
		select {
		case x := <-ch:
			fmt.Println("read channel value:", x)
		case ch <- i:
			fmt.Println("write channel value:", i)
		}
		time.Sleep(time.Second)
	}
}
