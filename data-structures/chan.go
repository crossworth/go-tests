package main

import (
	"fmt"
	"math/rand"
	"time"
)

func producer(c chan<- int) {
	for {
		n := rand.Intn(5000)
		fmt.Printf("Produced %v\n", n)
		c <- n
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	}
}

func consumer(ch <-chan int) {
	for i := range ch {
		time.Sleep(time.Duration(rand.Intn(5000)) * time.Millisecond)
		fmt.Printf("Consumed %v\n", i)
	}
}

func main() {
	ch := make(chan int)

	go producer(ch)
	go consumer(ch)
	go consumer(ch)
	go consumer(ch)
	consumer(ch)
	//for {
	//	select {
	//	case n := <-ch:
	//		fmt.Println(n)
	//	}
	//}

	//for n := range ch {
	//	fmt.Println(n)
	//}
}
