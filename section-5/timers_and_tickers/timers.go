package main

import (
	"fmt"
	"time"
)

func main() {
	nc := make(chan int)
	stopc := make(chan bool)

	go SlowCounter(1, nc, stopc)

	nc <- 2
	time.Sleep(6*time.Second)

	stopc <- true
	time.Sleep(1 * time.Second)
}

func SlowCounter(n int, nc chan int, stopc chan bool) {
	i := 0
	// create a duration of seconds
	d := time.Duration(n) * time.Second

	for {
		t := time.NewTimer(d)
		<- t.C
		i++
		fmt.Println(i)
	}
}
