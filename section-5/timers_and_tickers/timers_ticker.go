package main

import (
	"fmt"
	"time"
)

func main() {
	go tickCounter(2)
	time.Sleep(5 * time.Second)
}

func tickCounter(n int) {
	ticker := time.NewTicker(time.Duration(n) * time.Second)
	i := 0
	for t := range ticker.C {
		i++
		fmt.Println("Count ", i, " at ", t)

	}
}
